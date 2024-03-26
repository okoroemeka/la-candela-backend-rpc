package gapi

import (
	"context"
	"github.com/google/uuid"
	custom_error "github.com/okoroemeka/la-candela-backend-rpc/custom-error"
	"github.com/okoroemeka/la-candela-backend-rpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	uuidString := req.GetId()
	id, err := uuid.Parse(uuidString)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s, must be of type uuid", uuidString)
	}

	_, err = server.store.GetCourseById(ctx, id)

	if err != nil {
		if err.Error() == custom_error.NoRowFound {
			return nil, status.Errorf(codes.NotFound, "course does not exist")
		}
		return nil, status.Errorf(codes.Internal, "unable to retrieve course")
	}

	err = server.store.DeleteCourse(ctx, id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to delete course")
	}

	return &pb.DeleteCourseResponse{Message: "course deleted successfully"}, nil
}
