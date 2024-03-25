CREATE TABLE "verify_email" (
    id uuid UNIQUE PRIMARY KEY NOT NULL,
    email VARCHAR NOT NULL,
    secret_code VARCHAR NOT NULL,
    is_used VARCHAR NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    expired_at timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);