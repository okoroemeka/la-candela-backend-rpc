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

func (server *Server) UpdateCourse(ctx context.Context, req *pb.UpdateCourseRequest) (*pb.UpdateCourseResponse, error) {

	violations := validateReqParams(req)

	if violations != nil {
		return nil, custom_error.InvalidArgument(violations)
	}

	arg := getUpdateCourseParam(req)

	course, err := server.store.UpdateCourse(ctx, arg)

	if err != nil {
		if err.Error() == custom_error.NoRowFound {
			return nil, status.Errorf(codes.NotFound, "course does not exist")
		}
		return nil, status.Errorf(codes.Internal, "could not update course")
	}

	return &pb.UpdateCourseResponse{Course: ConvertCourse(course)}, nil
}

func validateReqParams(req *pb.UpdateCourseRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	locationId := req.GetLocationId()

	if err := validator.ValidateUuid(locationId); err != nil && len(locationId) > 0 {
		violations = append(violations, custom_error.FieldValidation("location_id", err))
	}

	priceHikeId := req.GetPriceHikeId()
	if err := validator.ValidateUuid(priceHikeId); err != nil && len(priceHikeId) > 0 {
		violations = append(violations, custom_error.FieldValidation("price_hike_id", err))
	}

	if err := validator.ValidateUuid(req.GetId()); err != nil {
		violations = append(violations, custom_error.FieldValidation("id", err))
	}

	courseLevel := req.GetCourseLevel()

	if err := validator.ValidateCourseLevel(courseLevel); err != nil && len(courseLevel) > 0 {
		violations = append(violations, custom_error.FieldValidation("course_level", err))
	}

	if err := validator.ValidateCourseDate(req.GetStartDate()); err != nil && req.GetEndDate() != nil {
		violations = append(violations, custom_error.FieldValidation("start_date", err))
	}

	if err := validator.ValidateCourseDate(req.GetEndDate()); err != nil && req.GetEndDate() != nil {
		violations = append(violations, custom_error.FieldValidation("end_date", err))
	}

	return
}

func getUpdateCourseParam(req *pb.UpdateCourseRequest) db.UpdateCourseParams {
	courseId, _ := uuid.Parse(req.GetId())

	priceHikeId, _ := uuid.Parse(req.GetPriceHikeId())
	locationId, _ := uuid.Parse(req.GetLocationId())

	courseTitle := req.GetCourseTitle()
	courseDesc := req.GetCourseDescription()
	price := req.GetPrice()

	return db.UpdateCourseParams{
		CourseTitle: pgtype.Text{
			String: courseTitle,
			Valid:  len(courseTitle) > 0,
		},

		CourseDescription: pgtype.Text{
			String: courseDesc,
			Valid:  len(courseDesc) > 0,
		},

		CourseLevel: db.NullCourseLevel{
			CourseLevel: db.CourseLevel(req.GetCourseLevel()),
			Valid:       true,
		},

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
		Price: pgtype.Int8{
			Int64: price,
			Valid: price > 0,
		},
		LocationID: pgtype.UUID{
			Bytes: locationId,
			Valid: true,
		},
		MinCapacity: pgtype.Int8{
			Int64: int64(req.GetMinCapacity()),
			Valid: true,
		},
		Open: pgtype.Bool{
			Bool:  req.GetOpen(),
			Valid: true,
		},
		PriceHikeID: pgtype.UUID{
			Bytes: priceHikeId,
			Valid: true,
		},
		ID: courseId,
	}
}
