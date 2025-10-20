-- name: CreateCustomer :one
INSERT INTO "customers" ("name", "email", "phone", "birthdate", "cpf")
VALUES ($1, $2, $3, $4, $5)
RETURNING "id", "name", "email", "phone", "birthdate", "cpf";

-- name: UpdateCustomer :one
UPDATE "customers"
SET
  "name" = sqlc.arg('name'),
  "email" = sqlc.arg('email'),
  "phone" = sqlc.arg('phone'),
  "birthdate" = sqlc.arg('birthdate'),
  "cpf" = sqlc.arg('cpf')
WHERE "id" = sqlc.arg('id')
RETURNING "id", "name", "email", "phone", "birthdate", "cpf";

-- name: GetCustomerByID :one
SELECT "id", "name", "email", "phone", "birthdate", "cpf"
FROM "customers"
WHERE "id" = $1
LIMIT 1;

-- name: GetCustomerByPhone :one
SELECT "id", "name", "email", "phone", "birthdate", "cpf"
FROM "customers"
WHERE "phone" = $1
LIMIT 1;

-- name: ListCustomers :many
SELECT "id", "name", "email", "phone", "birthdate", "cpf"
FROM "customers"
WHERE true
  AND (sqlc.arg('name')::text = '' OR "name" ILIKE '%' || sqlc.arg('name') || '%')
  AND (sqlc.arg('cpf')::text = '' OR "cpf" = sqlc.arg('cpf'))
  AND (sqlc.arg('phone')::text = '' OR "phone" = sqlc.arg('phone'))
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: CountCustomers :one
SELECT COUNT(id)
FROM "customers"
WHERE true
  AND (sqlc.arg('name')::text = '' OR "name" ILIKE '%' || sqlc.arg('name') || '%')
  AND (sqlc.arg('cpf')::text = '' OR "cpf" = sqlc.arg('cpf'))
  AND (sqlc.arg('phone')::text = '' OR "phone" = sqlc.arg('phone'));

-- name: DeleteCustomerByID :execrows
DELETE FROM "customers"
WHERE "id" = $1;
