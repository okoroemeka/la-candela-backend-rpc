-- noinspection SqlNoDataSourceInspectionForFile

CREATE TYPE "role" AS ENUM (
    'admin',
    'teacher',
    'student'
    );

CREATE TYPE "course_level" AS ENUM (
    'beginner',
    'intermediate',
    'advance'
    );

CREATE TYPE "course_role" AS ENUM (
    'leader',
    'follower'
    );

CREATE TABLE "user" (
                        "id" uuid UNIQUE PRIMARY KEY NOT NULL,
                        "first_name" varchar NOT NULL,
                        "last_name" varchar NOT NULL,
                        "email" varchar UNIQUE NOT NULL,
                        "is_verified" boolean NOT NULL DEFAULT false,
                        "password" varchar  NOT NULL,
                        password_changed_at timestamptz ,
                        "user_role" role NOT NULL DEFAULT 'student',
                        "created_at" timestamptz NOT NULL DEFAULT (now()),
                        "updated_at" timestamptz
);

CREATE TABLE "course" (
                          "id" uuid UNIQUE PRIMARY KEY NOT NULL,
                          "course_title" varchar NOT NULL,
                          "course_description" varchar NOT NULL,
                          "course_level" course_level NOT NULL DEFAULT 'beginner',
                          "start_date" timestamptz NOT NULL DEFAULT (now()),
                          "end_date" timestamptz NOT NULL DEFAULT (now()),
                          "start_time" time NOT NULL,
                          "end_time" time NOT NULL,
                          "price" bigint NOT NULL DEFAULT 0,
                          "location_id" uuid NOT NULL,
                          "min_capacity" bigint NOT NULL DEFAULT 0,
                          "open" boolean NOT NULL DEFAULT false,
                          "price_hike_id" uuid NOT NULL,
                          "created_by" uuid NOT NULL,
                          "created_at" timestamptz NOT NULL DEFAULT (now()),
                          "updated_at" timestamptz
);

CREATE TABLE "course_teachers" (
                                   "id" uuid UNIQUE PRIMARY KEY NOT NULL,
                                   "course_id" uuid NOT NULL,
                                   "teacher_id" uuid NOT NULL,
                                   "created_at" timestamptz NOT NULL DEFAULT (now()),
                                   "updated_at" timestamptz
);

CREATE TABLE "course_discount" (
                                   "id" uuid UNIQUE PRIMARY KEY NOT NULL,
                                   "course_id" uuid NOT NULL,
                                   "discount_id" uuid NOT NULL,
                                   "created_at" timestamptz NOT NULL DEFAULT (now()),
                                   "updated_at" timestamptz
);

CREATE TABLE "teacher" (
                           "id" uuid UNIQUE PRIMARY KEY NOT NULL,
                           "user_id" uuid NOT NULL,
                           "teachers_story" varchar DEFAULT null,
                           "created_at" timestamptz NOT NULL DEFAULT (now()),
                           "updated_at" timestamptz
);

CREATE TABLE "discount" (
                            "id" uuid UNIQUE PRIMARY KEY NOT NULL,
                            "title" varchar NOT NULL,
                            "description" varchar DEFAULT null,
                            "percentage" bigint DEFAULT 0,
                            "created_at" timestamptz NOT NULL DEFAULT (now()),
                            "updated_at" timestamptz
);

CREATE TABLE "location" (
                            "id" uuid UNIQUE PRIMARY KEY NOT NULL,
                            "street" varchar NOT NULL,
                            "house_num" bigint NOT NULL,
                            "city" varchar NOT NULL,
                            "zip_code" integer NOT NULL,
                            "created_at" timestamptz NOT NULL DEFAULT (now()),
                            "updated_at" timestamptz
);

CREATE TABLE "enrolment" (
                             "id" uuid UNIQUE PRIMARY KEY NOT NULL,
                             "course_id" uuid NOT NULL,
                             "user_id" uuid NOT NULL,
                             "course_role" course_role NOT NULL,
                             "paid" boolean NOT NULL,
                             "confirmed" boolean NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT (now()),
                             "updated_at" timestamptz
);

CREATE TABLE "price_hike" (
                              "id" uuid UNIQUE PRIMARY KEY NOT NULL,
                              "percentage" bigint DEFAULT 0,
                              "created_at" timestamptz NOT NULL DEFAULT (now()),
                              "updated_at" timestamptz
);

ALTER TABLE "course" ADD FOREIGN KEY ("location_id") REFERENCES "location" ("id");

ALTER TABLE "course" ADD FOREIGN KEY ("price_hike_id") REFERENCES "price_hike" ("id");

ALTER TABLE "course" ADD FOREIGN KEY ("created_by") REFERENCES "user" ("id");

ALTER TABLE "course_teachers" ADD FOREIGN KEY ("course_id") REFERENCES "course" ("id");

ALTER TABLE "course_teachers" ADD FOREIGN KEY ("teacher_id") REFERENCES "teacher" ("id");

ALTER TABLE "course_discount" ADD FOREIGN KEY ("course_id") REFERENCES "course" ("id");

ALTER TABLE "course_discount" ADD FOREIGN KEY ("discount_id") REFERENCES "discount" ("id");

ALTER TABLE "teacher" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "enrolment" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

CREATE UNIQUE INDEX ON "enrolment" ("course_id", "user_id");
CREATE UNIQUE INDEX ON "course_teachers" ("course_id", "teacher_id");
CREATE UNIQUE INDEX ON "course_discount" ("course_id", "discount_id");
CREATE INDEX ON "course" ("id", "start_date","course_level");