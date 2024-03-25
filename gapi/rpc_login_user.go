package gapi

import (
	"context"
	"database/sql"
	"errors"
	custom_error "github.com/okoroemeka/la-candela-backend-rpc/custom-error"
	"github.com/okoroemeka/la-candela-backend-rpc/pb"
	"github.com/okoroemeka/la-candela-backend-rpc/util"
	"github.com/okoroemeka/la-candela-backend-rpc/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	violations := validateLogin(req)

	if violations != nil {
		return nil, custom_error.InvalidArgument(violations)
	}

	user, err := server.store.GetUser(ctx, req.GetEmail())

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "please sign up to continue")
		}

		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	err = util.ComparePassword(req.GetPassword(), user.Password)

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "invalid password")
	}

	resp := &pb.LoginUserResponse{User: ConvertUser(user)}

	return resp, nil
}

func validateLogin(req *pb.LoginUserRequest) (violation []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateEmail(req.GetEmail()); err != nil {
		violation = append(violation, custom_error.FieldValidation(req.GetEmail(), err))
	}
	if err := validator.ValidatePassword(req.GetPassword()); err != nil {
		violation = append(violation, custom_error.FieldValidation("password", err))
	}
	return violation
}
