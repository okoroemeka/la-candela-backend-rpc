package gapi

import (
	db "github.com/okoroemeka/la-candela-backend-rpc/db/sqlc"
	"github.com/okoroemeka/la-candela-backend-rpc/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertUser(user db.User) *pb.User {
	return &pb.User{
		Id:                user.ID.String(),
		Email:             user.Email,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt.Time),
		CreatedAt:         timestamppb.New(user.CreatedAt),
		IsEmailVerified:   user.IsVerified,
	}
}

func ConvertCourse(course db.Course) *pb.Course {
	return &pb.Course{
		Id:          course.ID.String(),
		CourseTitle: course.CourseTitle,
		CourseLevel: string(course.CourseLevel),
		StartDate:   timestamppb.New(course.StartDate),
		EndDate:     timestamppb.New(course.EndDate),
		StartTime: &timestamppb.Timestamp{
			Seconds: course.StartTime.Microseconds,
		},
		EndTime: &timestamppb.Timestamp{
			Seconds: course.EndTime.Microseconds,
		},
		Price:             course.Price,
		LocationId:        course.LocationID.String(),
		MinCapacity:       int32(course.MinCapacity),
		Open:              false,
		PriceHikeId:       course.PriceHikeID.String(),
		CreatedAt:         timestamppb.New(course.CreatedAt),
		UpdatedAt:         timestamppb.New(course.UpdatedAt.Time),
		CreatedBy:         course.CreatedBy.String(),
		CourseDescription: course.CourseDescription,
	}
}
