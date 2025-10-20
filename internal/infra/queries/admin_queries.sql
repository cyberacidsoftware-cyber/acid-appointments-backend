-- name: CreateAdmin :one
INSERT INTO "admins" ("name", "email", "password")
VALUES ($1, $2, $3)
RETURNING "id";

