-- name: CreateSecretary :one
INSERT INTO "secretaries" ("name", "email", "password", "phone", "birthdate", "cpf", "cnpj")
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING "id", "name", "email", "password", "phone", "birthdate", "cpf", "cnpj";

-- name: UpdateSecretary :one
UPDATE "secretaries"
SET
  "name" = sqlc.arg('name'),
  "email" = sqlc.arg('email'),
  "password" = sqlc.arg('password'),
  "phone" = sqlc.arg('phone'),
  "birthdate" = sqlc.arg('birthdate'),
  "cpf" = sqlc.arg('cpf'),
  "cnpj" = sqlc.arg('cnpj')
WHERE "id" = sqlc.arg('id')
RETURNING "id", "name", "email", "password", "phone", "birthdate", "cpf", "cnpj";

-- name: GetSecretaryByID :one
SELECT "id", "name", "email", "password", "phone", "birthdate", "cpf", "cnpj"
FROM "secretaries"
WHERE "id" = $1
LIMIT 1;

-- name: GetSecretaryByEmail :one
SELECT "id", "name", "email", "password", "phone", "birthdate", "cpf", "cnpj"
FROM "secretaries"
WHERE "email" = $1
LIMIT 1;

-- name: ListSecretaries :many
SELECT "id", "name", "email", "password", "phone", "birthdate", "cpf", "cnpj"
FROM "secretaries"
WHERE true
  AND (sqlc.arg('name')::text = '' OR "name" ILIKE '%' || sqlc.arg('name') || '%')
  AND (sqlc.arg('cpf')::text = '' OR "cpf" = sqlc.arg('cpf'))
  AND (sqlc.arg('phone')::text = '' OR "phone" = sqlc.arg('phone'))
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: CountSecretaries :one
SELECT COUNT(id)
FROM "secretaries"
WHERE true
  AND (sqlc.arg('name')::text = '' OR "name" ILIKE '%' || sqlc.arg('name') || '%')
  AND (sqlc.arg('cpf')::text = '' OR "cpf" = sqlc.arg('cpf'))
  AND (sqlc.arg('phone')::text = '' OR "phone" = sqlc.arg('phone'));

-- name: DeleteSecretaryByID :execrows
DELETE FROM "secretaries"
WHERE "id" = $1;
