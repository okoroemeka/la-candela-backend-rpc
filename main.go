package main

import (
	"context"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/okoroemeka/la-candela-backend-rpc/db/sqlc"
	_ "github.com/okoroemeka/la-candela-backend-rpc/doc/statik"
	"github.com/okoroemeka/la-candela-backend-rpc/gapi"
	"github.com/okoroemeka/la-candela-backend-rpc/mail"
	"github.com/okoroemeka/la-candela-backend-rpc/pb"
	"github.com/okoroemeka/la-candela-backend-rpc/util"
	"github.com/okoroemeka/la-candela-backend-rpc/worker"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
	"os"
)

func main() {
	config, err := util.LoadConfig(".")

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if err != nil {
		log.Fatal().Err(err).Msg("Could not load config variables")
		return
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot connect to database")
	}

	runDBMigration(config.MigrationUrl, config.DBSource)

	store := db.NewStore(connPool)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	go runGatewayServer(config, store, taskDistributor)
	go runTaskProcessor(config, redisOpt, store)
	runGrpcServer(config, store, taskDistributor)

}

func runTaskProcessor(config util.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.GmailUsername, config.GmailPassword)
	processor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("Task processor is running")

	if err := processor.Start(); err != nil {
		log.Fatal().Err(err).Msg("Cannot start task processor")
	}
}

func runGrpcServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create grpc server")
	}

	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)

	pb.RegisterLaCandelaBackendRPCServer(grpcServer, server)

	/*  Register reflection service on gRPC server
	this allows the client to know what services are available on the server
	and how to call them.
	*/
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create grpc listener")
	}

	log.Info().Msgf("grpc server is running on %s", listener.Addr().String())

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start grpc server")
	}
}

func runDBMigration(migrationUrl, dbSource string) {
	migration, err := migrate.New(migrationUrl, dbSource)

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create migration instance:")
	}
	if err := migration.Up(); !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal().Err(err).Msg("Cannot run migration")
	}

	log.Info().Msg("Migration ran successfully")
}

func runGatewayServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create server")
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	jsonOptions := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOptions)
	err = pb.RegisterLaCandelaBackendRPCHandlerServer(ctx, grpcMux, server)

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot register gateway server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create statik file system")
	}

	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))

	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot create http listener")
	}

	log.Info().Msgf("gateway server is running on %s", listener.Addr().String())

	handler := gapi.HttpLogger(mux)
	err = http.Serve(listener, handler)

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot start http gateway server")
	}
}
