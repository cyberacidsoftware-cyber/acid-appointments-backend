-- name: CreateSpecialist :one
INSERT INTO specialists ("name", "email", "phone", "birthdate", "cpf", "cnpj")
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING "id";

-- name: CountSpecialists :one
SELECT COUNT("s"."id")
FROM "specialists" "s";

-- name: UpdateSpecialist :one
UPDATE "specialists"
SET 
  "name" = sqlc.arg('name'),
  "email" = sqlc.arg('email'),
  "phone" = sqlc.arg('phone'),
  "birthdate" = sqlc.arg('birthdate'),
  "cpf" = sqlc.arg('cpf'),
  "cnpj" = sqlc.arg('cnpj')
WHERE "id" = sqlc.arg('id')
RETURNING "id", "name", "email", "phone", "birthdate", "cpf", "cnpj";

-- name: GetSpecialistByID :one
SELECT "id", "name", "email", "phone", "birthdate", "cpf", "cnpj"
FROM "specialists"
WHERE "id" = $1
LIMIT 1;

-- name: GetSpecialistByEmail :one
SELECT "id", "name", "email", "phone", "birthdate", "cpf", "cnpj"
FROM "specialists"
WHERE "email" = $1
LIMIT 1;

-- name: ListSpecialists :many
SELECT "id", "name", "email", "phone", "birthdate", "cpf", "cnpj"
FROM "specialists"
LIMIT $1
OFFSET $2;

-- name: DeleteSpecialistByID :execrows
DELETE FROM "specialists"
WHERE "id" = $1;

-- name: ListServicesBySpecialistID :many
SELECT "s"."id",
    "sn"."specialization_id",
    "sn"."id" as "service_name_id",
    "sn"."name" as "service_name",
    "s"."price", "s"."duration"
FROM "services" "s"
JOIN "service_names" "sn" ON "sn"."id" = "s"."service_name_id"
WHERE "s"."specialist_id" = $1;
