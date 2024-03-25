package custom_error

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	UniqueViolation     = "23505"
	ForeignKeyViolation = "25503"
)

func ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return ""
	}

	return pgErr.Code
}

func FieldValidation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func InvalidArgument(violations []*errdetails.BadRequest_FieldViolation) (validationError error) {
	badRequest := &errdetails.BadRequest{FieldViolations: violations}
	statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")

	statusDetails, err := statusInvalid.WithDetails(badRequest)

	if err != nil {
		validationError = statusInvalid.Err()
		return
	}
	validationError = statusDetails.Err()

	return validationError
}
