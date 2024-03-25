package validator

import (
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/mail"
	"regexp"
	"time"
)

var (
	isValidName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, min, max int) error {
	valueLength := len(value)

	if valueLength < min || valueLength > max {
		return fmt.Errorf("value must be from %d-%d", min, max)
	}
	return nil
}

func ValidateName(value string) error {
	if !isValidName(value) {
		return fmt.Errorf("invalid value: %s, must contain only letters and/or spaces", value)
	}
	return ValidateString(value, 1, 20)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 5, 100); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("value must be a v")
	}
	return nil
}
func ValidatePassword(value string) error {
	return ValidateString(value, 5, 100)
}

const (
	BEGINNER     = "beginner"
	INTERMEDIATE = "intermediate"
	ADVANCE      = "advance"
)

func ValidateCourseLevel(courseLevel string) error {
	formattedCourseLevel := string(courseLevel)
	if (formattedCourseLevel != BEGINNER) && (formattedCourseLevel != INTERMEDIATE) && (formattedCourseLevel != ADVANCE) {
		return fmt.Errorf("invalid value: %s must be %s, %s or %s", courseLevel, BEGINNER, INTERMEDIATE, ADVANCE)
	}
	return nil
}

func ValidatePositiveInt(value int64) error {
	if value < 0 {
		return fmt.Errorf("invalid value: %d, it must be a positive integer", value)
	}
	return nil
}

func ValidateCourseDate(date *timestamppb.Timestamp) error {
	now := time.Now().Unix()

	//fmt.Printf("now %d", now)
	fmt.Println(GetFutureDateInUnix("12-1-2029", "02-1-2006"))

	if now >= date.Seconds {
		return fmt.Errorf("invalid value: %s, date must be in the future", date)
	}

	return nil
}

func ValidateUuid(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid value: %s, %s must be of type uuid", id, id)
	}
	return nil
}
func GetFutureDateInUnix(dateStr, layout string) (int64, error) {
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0, fmt.Errorf("error formatting date: %s", dateStr)
	}
	return date.Unix(), nil
}
