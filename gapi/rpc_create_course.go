package gapi

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	custom_error "github.com/okoroemeka/la-candela-backend-rpc/custom-error"
	db "github.com/okoroemeka/la-candela-backend-rpc/db/sqlc"
	"github.com/okoroemeka/la-candela-backend-rpc/pb"
	"github.com/okoroemeka/la-candela-backend-rpc/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CreateCourseResponse, error) {
	violations := ValidateCreateCourseRequest(req)

	if violations != nil {
		return nil, custom_error.InvalidArgument(violations)
	}

	courseArg := getCreateCourseParam(req)

	course, err := server.store.CreateCourse(ctx, courseArg)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not create course")
	}

	resp := &pb.CreateCourseResponse{Course: ConvertCourse(course)}

	return resp, nil
}

func ValidateCreateCourseRequest(req *pb.CreateCourseRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateString(req.GetCourseTitle(), 1, 250); err != nil {
		violations = append(violations, custom_error.FieldValidation("course_title", err))
	}
	if err := validator.ValidateString(req.GetCourseDescription(), 1, 550); err != nil {
		violations = append(violations, custom_error.FieldValidation("course_description", err))
	}
	if err := validator.ValidateCourseLevel(req.GetCourseLevel()); err != nil {
		violations = append(violations, custom_error.FieldValidation("course_level", err))
	}
	if err := validator.ValidatePositiveInt(req.GetPrice()); err != nil {
		violations = append(violations, custom_error.FieldValidation("price", err))
	}
	if err := validator.ValidateString(req.GetLocationId(), 36, 36); err != nil {
		violations = append(violations, custom_error.FieldValidation("location_id", err))
	}
	if err := validator.ValidatePositiveInt(int64(req.GetMinCapacity())); err != nil {
		violations = append(violations, custom_error.FieldValidation("min_capacity", err))
	}
	if err := validator.ValidateUuid(req.GetPriceHikeId()); err != nil {
		violations = append(violations, custom_error.FieldValidation("price_hike_id", err))
	}
	if err := validator.ValidateUuid(req.GetLocationId()); err != nil {
		violations = append(violations, custom_error.FieldValidation("location_id", err))
	}
	if err := validator.ValidateUuid(req.GetCreatedBy()); err != nil {
		violations = append(violations, custom_error.FieldValidation("created_by", err))
	}
	if err := validator.ValidateCourseDate(req.GetStartDate()); err != nil {
		violations = append(violations, custom_error.FieldValidation("start_date", err))
	}
	if err := validator.ValidateCourseDate(req.GetEndDate()); err != nil {
		violations = append(violations, custom_error.FieldValidation("end_date", err))
	}
	return
}

func getCreateCourseParam(req *pb.CreateCourseRequest) db.CreateCourseParams {
	locationId, _ := uuid.Parse(req.GetLocationId())
	priceHikeId, _ := uuid.Parse(req.GetPriceHikeId())
	createdBy, _ := uuid.Parse(req.GetCreatedBy())

	return db.CreateCourseParams{
		ID:                uuid.New(),
		CourseTitle:       req.GetCourseTitle(),
		CourseDescription: req.GetCourseDescription(),
		CourseLevel:       db.CourseLevel(req.GetCourseLevel()),
		StartDate: pgtype.Timestamptz{
			Time:  req.GetStartDate().AsTime(),
			Valid: true,
		},
		EndDate: pgtype.Timestamptz{
			Time:  req.GetEndDate().AsTime(),
			Valid: true,
		},
		StartTime: pgtype.Time{
			Microseconds: req.GetStartTime().GetSeconds(),
			Valid:        true,
		},
		EndTime: pgtype.Time{
			Microseconds: req.GetEndTime().GetSeconds(),
			Valid:        true,
		},
		Price:       req.GetPrice(),
		LocationID:  locationId,
		MinCapacity: int64(req.GetMinCapacity()),
		Open:        req.GetOpen(),
		PriceHikeID: priceHikeId,
		CreatedBy:   createdBy,
	}
}
