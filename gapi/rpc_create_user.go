package gapi

import (
	"context"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	custom_error "github.com/okoroemeka/la-candela-backend-rpc/custom-error"
	db "github.com/okoroemeka/la-candela-backend-rpc/db/sqlc"
	"github.com/okoroemeka/la-candela-backend-rpc/pb"
	"github.com/okoroemeka/la-candela-backend-rpc/util"
	"github.com/okoroemeka/la-candela-backend-rpc/validator"
	"github.com/okoroemeka/la-candela-backend-rpc/worker"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := ValidateCreateUserRequest(req)

	if violations != nil {
		return nil, custom_error.InvalidArgument(violations)
	}

	hashedPass, err := util.HashPassword(req.GetPassword())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot hash password: %s", err)
	}

	user := db.CreateUserParams{
		ID:        uuid.New(),
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
		Password:  hashedPass,
	}

	txArg := db.CreateUserTxParams{
		CreateUserParams: user,
		AfterCreate: func(user db.User) error {
			opts := []asynq.Option{
				asynq.MaxRetry(10),
				asynq.ProcessIn(10 * time.Second),
				asynq.Queue(worker.QueCritical),
			}
			return server.taskDistributor.DistributeTaskSendVerifyEmail(ctx, &worker.PayloadSendVerifyEmail{
				UserID: user.ID,
				Email:  user.Email,
			}, opts...)
		},
	}

	txResult, err := server.store.CreateUserTransaction(ctx, txArg)

	if err != nil {
		if custom_error.ErrorCode(err) == custom_error.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Errorf(codes.Internal, "could not create user")
	}

	resp := &pb.CreateUserResponse{User: ConvertUser(txResult.User)}

	return resp, nil
}

func ValidateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateName(req.GetFirstName()); err != nil {
		violations = append(violations, custom_error.FieldValidation("first_name", err))
	}
	if err := validator.ValidateName(req.GetLastName()); err != nil {
		violations = append(violations, custom_error.FieldValidation("last_name", err))
	}
	if err := validator.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, custom_error.FieldValidation("email", err))
	}
	if err := validator.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, custom_error.FieldValidation("password", err))
	}

	return
}
