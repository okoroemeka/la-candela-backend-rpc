package gapi

import (
	"fmt"
	db "github.com/okoroemeka/la-candela-backend-rpc/db/sqlc"
	"github.com/okoroemeka/la-candela-backend-rpc/pb"
	"github.com/okoroemeka/la-candela-backend-rpc/token"
	"github.com/okoroemeka/la-candela-backend-rpc/util"
	"github.com/okoroemeka/la-candela-backend-rpc/worker"
)

type Server struct {
	pb.UnimplementedLaCandelaBackendRPCServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
