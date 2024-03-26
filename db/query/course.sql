-- name: CreateCourse :one
INSERT INTO "course"
    ("id", "course_title","course_description","course_level","start_date","end_date","start_time","end_time","price","location_id","min_capacity","open","price_hike_id","created_by")
VALUES
    ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)
RETURNING *;

-- name: GetCourseById :one
SELECT * FROM
    "course"
WHERE
    id = sqlc.arg(id)
ORDER BY
    "course_title"
LIMIT 1;

-- name: GetAllCourse :many
SELECT * FROM "course" ORDER BY "created_at";

-- name: UpdateCourse :one
UPDATE "course"
SET
    course_title = coalesce(sqlc.narg(course_title),course_title),
    course_description = coalesce(sqlc.narg(course_description),course_description),
    course_level = coalesce(sqlc.narg(course_level),course_level),
    start_date = coalesce(sqlc.narg(start_date), start_date),
    end_date = coalesce(sqlc.narg(end_date), end_date),
    start_time = coalesce(sqlc.narg(start_time), start_time),
    end_time = coalesce(sqlc.narg(end_time), end_time),
    price  = coalesce(sqlc.narg(price), price),
    location_id = coalesce(sqlc.narg(location_id), location_id),
    min_capacity = coalesce(sqlc.narg(min_capacity), min_capacity),
    open = coalesce(sqlc.narg(open), open),
    price_hike_id = coalesce(sqlc.narg(price_hike_id), price_hike_id)
WHERE
    id = sqlc.arg(id)
RETURNING *;

-- name: DeleteCourse :exec
    DELETE FROM course WHERE id = sqlc.arg(id);

--     created_by = coalesce(sqlc.narg(created_by), created_by)
