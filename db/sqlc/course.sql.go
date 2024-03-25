// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: course.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createCourse = `-- name: CreateCourse :one
INSERT INTO "course"
    ("id", "course_title","course_description","course_level","start_date","end_date","start_time","end_time","price","location_id","min_capacity","open","price_hike_id","created_by")
VALUES
    ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)
RETURNING id, course_title, course_description, course_level, start_date, end_date, start_time, end_time, price, location_id, min_capacity, open, price_hike_id, created_by, created_at, updated_at
`

type CreateCourseParams struct {
	ID                uuid.UUID          `json:"id"`
	CourseTitle       string             `json:"course_title"`
	CourseDescription string             `json:"course_description"`
	CourseLevel       CourseLevel        `json:"course_level"`
	StartDate         pgtype.Timestamptz `json:"start_date"`
	EndDate           pgtype.Timestamptz `json:"end_date"`
	StartTime         pgtype.Time        `json:"start_time"`
	EndTime           pgtype.Time        `json:"end_time"`
	Price             int64              `json:"price"`
	LocationID        uuid.UUID          `json:"location_id"`
	MinCapacity       int64              `json:"min_capacity"`
	Open              bool               `json:"open"`
	PriceHikeID       uuid.UUID          `json:"price_hike_id"`
	CreatedBy         uuid.UUID          `json:"created_by"`
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) (Course, error) {
	row := q.db.QueryRow(ctx, createCourse,
		arg.ID,
		arg.CourseTitle,
		arg.CourseDescription,
		arg.CourseLevel,
		arg.StartDate,
		arg.EndDate,
		arg.StartTime,
		arg.EndTime,
		arg.Price,
		arg.LocationID,
		arg.MinCapacity,
		arg.Open,
		arg.PriceHikeID,
		arg.CreatedBy,
	)
	var i Course
	err := row.Scan(
		&i.ID,
		&i.CourseTitle,
		&i.CourseDescription,
		&i.CourseLevel,
		&i.StartDate,
		&i.EndDate,
		&i.StartTime,
		&i.EndTime,
		&i.Price,
		&i.LocationID,
		&i.MinCapacity,
		&i.Open,
		&i.PriceHikeID,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCourse = `-- name: DeleteCourse :exec
    DELETE FROM course WHERE id = $1
`

func (q *Queries) DeleteCourse(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteCourse, id)
	return err
}

const getAllCourse = `-- name: GetAllCourse :many
SELECT id, course_title, course_description, course_level, start_date, end_date, start_time, end_time, price, location_id, min_capacity, open, price_hike_id, created_by, created_at, updated_at FROM "course" ORDER BY "created_at"
`

func (q *Queries) GetAllCourse(ctx context.Context) ([]Course, error) {
	rows, err := q.db.Query(ctx, getAllCourse)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Course{}
	for rows.Next() {
		var i Course
		if err := rows.Scan(
			&i.ID,
			&i.CourseTitle,
			&i.CourseDescription,
			&i.CourseLevel,
			&i.StartDate,
			&i.EndDate,
			&i.StartTime,
			&i.EndTime,
			&i.Price,
			&i.LocationID,
			&i.MinCapacity,
			&i.Open,
			&i.PriceHikeID,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCourseById = `-- name: GetCourseById :one
SELECT id, course_title, course_description, course_level, start_date, end_date, start_time, end_time, price, location_id, min_capacity, open, price_hike_id, created_by, created_at, updated_at FROM
    "course"
WHERE
    id = $1
ORDER BY
    "course_title"
LIMIT 1
`

func (q *Queries) GetCourseById(ctx context.Context, id uuid.UUID) (Course, error) {
	row := q.db.QueryRow(ctx, getCourseById, id)
	var i Course
	err := row.Scan(
		&i.ID,
		&i.CourseTitle,
		&i.CourseDescription,
		&i.CourseLevel,
		&i.StartDate,
		&i.EndDate,
		&i.StartTime,
		&i.EndTime,
		&i.Price,
		&i.LocationID,
		&i.MinCapacity,
		&i.Open,
		&i.PriceHikeID,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCourse = `-- name: UpdateCourse :one
UPDATE "course"
SET
    course_title = coalesce($1,course_title),
    course_description = coalesce($2,course_description),
    course_level = coalesce($3,course_level),
    start_date = coalesce($4, start_date),
    end_date = coalesce($5, end_date),
    start_time = coalesce($6, start_time),
    end_time = coalesce($7, end_time),
    price  = coalesce($8, price),
    location_id = coalesce($9, location_id),
    min_capacity = coalesce($10, min_capacity),
    open = coalesce($11, open),
    price_hike_id = coalesce($12, price_hike_id),
    created_by = coalesce($13, created_by)
WHERE
    id = $14
RETURNING id, course_title, course_description, course_level, start_date, end_date, start_time, end_time, price, location_id, min_capacity, open, price_hike_id, created_by, created_at, updated_at
`

type UpdateCourseParams struct {
	CourseTitle       pgtype.Text        `json:"course_title"`
	CourseDescription pgtype.Text        `json:"course_description"`
	CourseLevel       NullCourseLevel    `json:"course_level"`
	StartDate         pgtype.Timestamptz `json:"start_date"`
	EndDate           pgtype.Timestamptz `json:"end_date"`
	StartTime         pgtype.Time        `json:"start_time"`
	EndTime           pgtype.Time        `json:"end_time"`
	Price             pgtype.Int8        `json:"price"`
	LocationID        pgtype.UUID        `json:"location_id"`
	MinCapacity       pgtype.Int8        `json:"min_capacity"`
	Open              pgtype.Bool        `json:"open"`
	PriceHikeID       pgtype.UUID        `json:"price_hike_id"`
	CreatedBy         pgtype.UUID        `json:"created_by"`
	ID                uuid.UUID          `json:"id"`
}

func (q *Queries) UpdateCourse(ctx context.Context, arg UpdateCourseParams) (Course, error) {
	row := q.db.QueryRow(ctx, updateCourse,
		arg.CourseTitle,
		arg.CourseDescription,
		arg.CourseLevel,
		arg.StartDate,
		arg.EndDate,
		arg.StartTime,
		arg.EndTime,
		arg.Price,
		arg.LocationID,
		arg.MinCapacity,
		arg.Open,
		arg.PriceHikeID,
		arg.CreatedBy,
		arg.ID,
	)
	var i Course
	err := row.Scan(
		&i.ID,
		&i.CourseTitle,
		&i.CourseDescription,
		&i.CourseLevel,
		&i.StartDate,
		&i.EndDate,
		&i.StartTime,
		&i.EndTime,
		&i.Price,
		&i.LocationID,
		&i.MinCapacity,
		&i.Open,
		&i.PriceHikeID,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
