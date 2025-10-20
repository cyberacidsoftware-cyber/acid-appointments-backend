-- name: CreateSpecialization :one
INSERT INTO "specializations" ("name")
VALUES ($1)
RETURNING "id";

-- name: UpdateSpecialization :one
UPDATE "specializations"
SET 
	"name" = sqlc.arg('name')
WHERE "id" = sqlc.arg('id')
RETURNING "id";

-- name: GetSpecializationByID :one
SELECT "id", "name"
FROM "specializations" 
WHERE "id" = $1;

-- name: GetSpecializationByName :one
SELECT "id", "name"
FROM "specializations" 
WHERE "name" = $1;

-- name: ListSpecializations :many
SELECT "id", "name"
FROM "specializations"
ORDER BY "name";

-- name: ListSpecializationsBySpecialistID :many
SELECT "sp"."id", "sp"."name"
FROM "specializations" "sp"
WHERE "sp"."id" IN (
	SELECT "sn"."specialization_id" 
	FROM "services" "s" 
	JOIN "service_names" "sn" ON "sn"."id" = "s"."service_name_id"
	WHERE "s"."specialist_id" = $1
);

-- name: DeleteSpecializationByID :execrows
DELETE FROM "specializations"
WHERE "id" = $1;
