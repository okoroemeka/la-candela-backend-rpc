// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user"
    ("id","first_name","last_name","email","password")
values($1,$2,$3,$4,$5) RETURNING id, first_name, last_name, email, is_verified, password, password_changed_at, user_role, created_at, updated_at
`

type CreateUserParams struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.IsVerified,
		&i.Password,
		&i.PasswordChangedAt,
		&i.UserRole,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, email, is_verified, password, password_changed_at, user_role, created_at, updated_at FROM "user" WHERE "email"=$1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.IsVerified,
		&i.Password,
		&i.PasswordChangedAt,
		&i.UserRole,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE "user"
SET
    first_name = coalesce($1, first_name),
    last_name = coalesce($2, last_name),
    email = coalesce($3, email),
    is_verified = coalesce($4, is_verified),
    password = coalesce($5, password),
    password_changed_at = coalesce($6, password_changed_at),
    user_role = coalesce($7, user_role),
    updated_at = coalesce($8,updated_at)
WHERE
    id = $9
RETURNING id, first_name, last_name, email, is_verified, password, password_changed_at, user_role, created_at, updated_at
`

type UpdateUserParams struct {
	FullName          pgtype.Text        `json:"full_name"`
	LastName          pgtype.Text        `json:"last_name"`
	Email             pgtype.Text        `json:"email"`
	IsVerified        pgtype.Bool        `json:"is_verified"`
	Password          pgtype.Text        `json:"password"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	UserRole          NullRole           `json:"user_role"`
	UpdatedAt         pgtype.Timestamptz `json:"updated_at"`
	ID                uuid.UUID          `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.FullName,
		arg.LastName,
		arg.Email,
		arg.IsVerified,
		arg.Password,
		arg.PasswordChangedAt,
		arg.UserRole,
		arg.UpdatedAt,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.IsVerified,
		&i.Password,
		&i.PasswordChangedAt,
		&i.UserRole,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const verifyUser = `-- name: VerifyUser :one
UPDATE "user" SET "is_verified"=$1 WHERE "id"=$2 RETURNING id, first_name, last_name, email, is_verified, password, password_changed_at, user_role, created_at, updated_at
`

type VerifyUserParams struct {
	IsVerified bool      `json:"is_verified"`
	ID         uuid.UUID `json:"id"`
}

func (q *Queries) VerifyUser(ctx context.Context, arg VerifyUserParams) (User, error) {
	row := q.db.QueryRow(ctx, verifyUser, arg.IsVerified, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.IsVerified,
		&i.Password,
		&i.PasswordChangedAt,
		&i.UserRole,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllUsers = `-- name: getAllUsers :many
SELECT id, first_name, last_name, email, is_verified, password, password_changed_at, user_role, created_at, updated_at FROM "user"
`

func (q *Queries) getAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.IsVerified,
			&i.Password,
			&i.PasswordChangedAt,
			&i.UserRole,
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
