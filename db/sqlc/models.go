// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type CourseLevel string

const (
	CourseLevelBeginner     CourseLevel = "beginner"
	CourseLevelIntermediate CourseLevel = "intermediate"
	CourseLevelAdvance      CourseLevel = "advance"
)

func (e *CourseLevel) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CourseLevel(s)
	case string:
		*e = CourseLevel(s)
	default:
		return fmt.Errorf("unsupported scan type for CourseLevel: %T", src)
	}
	return nil
}

type NullCourseLevel struct {
	CourseLevel CourseLevel `json:"course_level"`
	Valid       bool        `json:"valid"` // Valid is true if CourseLevel is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCourseLevel) Scan(value interface{}) error {
	if value == nil {
		ns.CourseLevel, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CourseLevel.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCourseLevel) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CourseLevel), nil
}

type CourseRole string

const (
	CourseRoleLeader   CourseRole = "leader"
	CourseRoleFollower CourseRole = "follower"
)

func (e *CourseRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CourseRole(s)
	case string:
		*e = CourseRole(s)
	default:
		return fmt.Errorf("unsupported scan type for CourseRole: %T", src)
	}
	return nil
}

type NullCourseRole struct {
	CourseRole CourseRole `json:"course_role"`
	Valid      bool       `json:"valid"` // Valid is true if CourseRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCourseRole) Scan(value interface{}) error {
	if value == nil {
		ns.CourseRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CourseRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCourseRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CourseRole), nil
}

type Role string

const (
	RoleAdmin   Role = "admin"
	RoleTeacher Role = "teacher"
	RoleStudent Role = "student"
)

func (e *Role) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Role(s)
	case string:
		*e = Role(s)
	default:
		return fmt.Errorf("unsupported scan type for Role: %T", src)
	}
	return nil
}

type NullRole struct {
	Role  Role `json:"role"`
	Valid bool `json:"valid"` // Valid is true if Role is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRole) Scan(value interface{}) error {
	if value == nil {
		ns.Role, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Role.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Role), nil
}

type Course struct {
	ID                uuid.UUID          `json:"id"`
	CourseTitle       string             `json:"course_title"`
	CourseDescription string             `json:"course_description"`
	CourseLevel       CourseLevel        `json:"course_level"`
	StartDate         time.Time          `json:"start_date"`
	EndDate           time.Time          `json:"end_date"`
	StartTime         pgtype.Time        `json:"start_time"`
	EndTime           pgtype.Time        `json:"end_time"`
	Price             int64              `json:"price"`
	LocationID        uuid.UUID          `json:"location_id"`
	MinCapacity       int64              `json:"min_capacity"`
	Open              bool               `json:"open"`
	PriceHikeID       uuid.UUID          `json:"price_hike_id"`
	CreatedBy         uuid.UUID          `json:"created_by"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         pgtype.Timestamptz `json:"updated_at"`
}

type CourseDiscount struct {
	ID         uuid.UUID          `json:"id"`
	CourseID   uuid.UUID          `json:"course_id"`
	DiscountID uuid.UUID          `json:"discount_id"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
}

type CourseTeacher struct {
	ID        uuid.UUID          `json:"id"`
	CourseID  uuid.UUID          `json:"course_id"`
	TeacherID uuid.UUID          `json:"teacher_id"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type Discount struct {
	ID          uuid.UUID          `json:"id"`
	Title       string             `json:"title"`
	Description pgtype.Text        `json:"description"`
	Percentage  pgtype.Int8        `json:"percentage"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

type Enrolment struct {
	ID         uuid.UUID          `json:"id"`
	CourseID   uuid.UUID          `json:"course_id"`
	UserID     uuid.UUID          `json:"user_id"`
	CourseRole CourseRole         `json:"course_role"`
	Paid       bool               `json:"paid"`
	Confirmed  bool               `json:"confirmed"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
}

type Location struct {
	ID        uuid.UUID          `json:"id"`
	Street    string             `json:"street"`
	HouseNum  int64              `json:"house_num"`
	City      string             `json:"city"`
	ZipCode   int32              `json:"zip_code"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type PriceHike struct {
	ID         uuid.UUID          `json:"id"`
	Percentage pgtype.Int8        `json:"percentage"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
}

type Teacher struct {
	ID            uuid.UUID          `json:"id"`
	UserID        uuid.UUID          `json:"user_id"`
	TeachersStory pgtype.Text        `json:"teachers_story"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
}

type User struct {
	ID                uuid.UUID          `json:"id"`
	FirstName         string             `json:"first_name"`
	LastName          string             `json:"last_name"`
	Email             string             `json:"email"`
	IsVerified        bool               `json:"is_verified"`
	Password          string             `json:"password"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	UserRole          Role               `json:"user_role"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         pgtype.Timestamptz `json:"updated_at"`
}

type VerifyEmail struct {
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email"`
	SecretCode string    `json:"secret_code"`
	IsUsed     string    `json:"is_used"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}
