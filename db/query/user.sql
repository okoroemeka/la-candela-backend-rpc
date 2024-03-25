-- name: CreateUser :one
INSERT INTO "user"
    ("id","first_name","last_name","email","password")
values($1,$2,$3,$4,$5) RETURNING *;

-- name: GetUser :one
SELECT * FROM "user" WHERE "email"=$1;

-- name: VerifyUser :one
UPDATE "user" SET "is_verified"=$1 WHERE "id"=$2 RETURNING *;

-- name: UpdateUser :one
UPDATE "user"
SET
    first_name = coalesce(sqlc.narg(full_name), first_name),
    last_name = coalesce(sqlc.narg(last_name), last_name),
    email = coalesce(sqlc.narg(email), email),
    is_verified = coalesce(sqlc.narg(is_verified), is_verified),
    password = coalesce(sqlc.narg(password), password),
    password_changed_at = coalesce(sqlc.narg(password_changed_at), password_changed_at),
    user_role = coalesce(sqlc.narg(user_role), user_role),
    updated_at = coalesce(sqlc.narg(updated_at),updated_at)
WHERE
    id = sqlc.arg(id)
RETURNING *;

-- name: getAllUsers :many
SELECT * FROM "user";


