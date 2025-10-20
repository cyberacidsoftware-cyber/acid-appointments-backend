-- name: CreateService :one
INSERT INTO services ("service_name_id", "specialist_id", "price", "duration")
VALUES ($1, $2, $3, $4)
RETURNING "id";

-- name: UpdateService :one
UPDATE "services" 
SET
    "price" = sqlc.arg('price'),
    "duration" = sqlc.arg('duration')
WHERE "id" = sqlc.arg('id')
RETURNING "id";

-- name: GetService :one
SELECT "id", "service_name_id", "specialist_id", "price", "duration"
FROM services
WHERE "service_name_id" = $1 AND "specialist_id" = $2;

-- name: GetServiceByID :one
SELECT "id", "service_name_id", "specialist_id", "price", "duration"
FROM services
WHERE "id" = $1;

-- name: DeleteService :execrows
DELETE FROM "services"
WHERE "id" = $1;
