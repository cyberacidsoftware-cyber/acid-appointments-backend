-- name: GetIdentityByEmail :one
SELECT "id", "name", "email", "password", 'admin' as "role"
FROM "admins"
WHERE "admins"."email" = $1
UNION
SELECT "id", "name", "email", "password", 'secretary' as "role"
FROM "secretaries"
WHERE "secretaries"."email" = $1
LIMIT 1;

-- name: GetIdentityByID :one
SELECT "id", "name", "email", "password", 'admin' as "role"
FROM "admins"
WHERE "admins"."id" = $1
UNION
SELECT "id", "name", "email", "password", 'secretary' as "role"
FROM "secretaries"
WHERE "secretaries"."id" = $1
LIMIT 1;
