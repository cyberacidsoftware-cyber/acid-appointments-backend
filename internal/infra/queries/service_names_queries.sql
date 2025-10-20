-- name: CreateServiceName :one
INSERT INTO "service_names" ("name", "specialization_id")
VALUES ($1, $2)
RETURNING "id";

-- name: UpdateServiceName :one
UPDATE "service_names"
SET
    "name" = sqlc.arg('name')
WHERE "id" = sqlc.arg('id')
RETURNING "id";

-- name: GetServiceNameByName :one
SELECT "id", "name", "specialization_id"
FROM "service_names" 
WHERE "name" = $1;

-- name: GetServiceNameByID :one
SELECT "id", "name", "specialization_id"
FROM "service_names" 
WHERE "id" = $1;

-- name: ListServiceNames :many
SELECT "id", "name", "specialization_id"
FROM "service_names"
ORDER BY "name";

-- name: DeleteServiceNameByID :execrows
DELETE FROM "service_names"
WHERE "id" = $1;