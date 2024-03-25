package gapi

import (
	"context"
	"github.com/google/uuid"
	custom_error "github.com/okoroemeka/la-candela-backend-rpc/custom-error"
	"github.com/okoroemeka/la-candela-backend-rpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.GetCourseResponse, error) {
	uuidString := req.GetId()

	id, err := uuid.Parse(uuidString)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s must be of type uuid", uuidString)
	}

	course, err := server.store.GetCourseById(ctx, id)
	
	if err != nil {
		if custom_error.ErrorCode(err) == custom_error.NoRowFound {
			return nil, status.Errorf(codes.NotFound, "course does not exist")
		}
		return nil, status.Errorf(codes.Internal, "could not retrieve course")
	}

	return &pb.GetCourseResponse{Course: ConvertCourse(course)}, nil
}
