-- name: CreateAppointment :one
INSERT INTO "appointments" ("customer_id", "specialist_id", "service_name_id", "price", "duration", "date", "time", "status")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING "id";

-- name: UpdateAppointment :one
UPDATE "appointments"
SET
  "date" = sqlc.arg('date'),
  "time" = sqlc.arg('time'),
  "status" = sqlc.arg('status')
WHERE "id" = sqlc.arg('id')
RETURNING "id", "customer_id", "specialist_id", "service_name_id", "price", "duration", "date", "time", "status", "notified_at", "notified_by";

-- name: ListAppointmentsBySpecialistID :many
SELECT "a"."id", "a"."price", "a"."duration", "a"."date", "a"."time", "a"."status", "a"."notified_at", "a"."notified_by",
  "a"."customer_id", "c"."name" as "customer_name",
  "a"."service_name_id", "sn"."name" as "service_name"
FROM "appointments" "a"
JOIN "customers" "c" ON "a"."customer_id" = "c"."id"
JOIN "service_names" "sn" ON "a"."service_name_id" = "sn"."id"
WHERE "a"."specialist_id" = $1 
  AND ($2::text = '' OR "a"."date" = DATE($2))
ORDER BY "a"."date" DESC, "a"."time" DESC;

-- name: GetAppointmentByID :one
SELECT "id", "customer_id", "specialist_id", "service_name_id", "price", "duration", "date", "time", "status", "notified_at", "notified_by"
FROM "appointments"
WHERE "id" = $1;

-- name: GetAppointmentEnrichedByID :one
SELECT "a"."id", "a"."price", "a"."duration", "a"."date", "a"."time", "a"."status", "a"."notified_at", "a"."notified_by",
  "a"."customer_id", "c"."name" as "customer_name",
  "a"."specialist_id", "s"."name" as "specialist_name",
  "a"."service_name_id", "sn"."name" as "service_name"
FROM "appointments" "a"
JOIN "customers" "c" ON "a"."customer_id" = "c"."id"
JOIN "specialists" "s" ON "a"."specialist_id" = "s"."id"
JOIN "service_names" "sn" ON "a"."service_name_id" = "sn"."id"
WHERE "a"."id" = $1;

-- name: ListAppointmentsByDate :many
SELECT "id", "customer_id", "specialist_id", "service_name_id", "price", "duration", "date", "time", "status", "notified_at", "notified_by"
FROM "appointments"
WHERE "date" = $1
ORDER BY "date" DESC, "time" DESC;

-- name: ListAppointments :many
SELECT "a"."id", "a"."price", "a"."duration", "a"."date", "a"."time", "a"."status", "a"."notified_at", "a"."notified_by",
  "a"."customer_id", "c"."name" as "customer_name",
  "a"."specialist_id", "s"."name" as "specialist_name",
  "a"."service_name_id", "sn"."name" as "service_name"
FROM "appointments" "a"
JOIN "customers" "c" ON "a"."customer_id" = "c"."id"
JOIN "specialists" "s" ON "a"."specialist_id" = "s"."id"
JOIN "service_names" "sn" ON "a"."service_name_id" = "sn"."id"
WHERE true
  AND ($1::date IS NULL OR "a"."date" >= $1) 
  AND ($2::date IS NULL OR "a"."date" <= $2)
  AND ($3::text = '' OR "c"."name" ILIKE '%' || $3 || '%')
  AND ($4::text = '' OR "s"."name" ILIKE '%' || $4 || '%')
  AND ($5::text = '' OR "sn"."name" ILIKE '%' || $5 || '%')
  AND ($6::int = 0 OR "a"."status" = $6)
ORDER BY "a"."date" DESC, "a"."time" DESC
LIMIT $7
OFFSET $8;

-- name: CountAppointments :one
SELECT COUNT("a"."id")
FROM "appointments" "a"
JOIN "customers" "c" ON "a"."customer_id" = "c"."id"
JOIN "specialists" "s" ON "a"."specialist_id" = "s"."id"
JOIN "service_names" "sn" ON "a"."service_name_id" = "sn"."id"
WHERE true
  AND ($1::date IS NULL OR "a"."date" >= $1) 
  AND ($2::date IS NULL OR "a"."date" <= $2)
  AND ($3::text = '' OR "c"."name" ILIKE '%' || $3 || '%')
  AND ($4::text = '' OR "s"."name" ILIKE '%' || $3 || '%')
  AND ($5::text = '' OR "sn"."name" ILIKE '%' || $3 || '%')
  AND ($6::int = 0 OR "a"."status" = $6);

-- name: ListAppointmentsCalendar :many
SELECT "a"."id", "a"."date", "a"."time", "a"."status", "s"."name" as "specialist_name"
FROM "appointments" "a"
JOIN "specialists" "s" ON "a"."specialist_id" = "s"."id"
WHERE true
  AND "a"."date" >= $1
  AND "a"."date" <= $2
ORDER BY "a"."date" DESC, "a"."time" DESC;

-- name: ListAppointmentsCalendarCount :many
SELECT date_part('month', "a"."date")::int as "month", "status", COUNT("a"."id")::int as "count"
FROM "appointments" "a"
WHERE "date" >= $1
  AND "date" <= $2
GROUP BY "month", "status"
ORDER BY "month" ASC;

-- name: AppointmentsIntersects :one
SELECT COUNT("date") > 0
FROM "appointments"
WHERE "date" = $1 
  AND "specialist_id" = $2
  AND (
    ("time" <= sqlc.arg('time')::time AND sqlc.arg('time')::time < "time" + "duration") OR
    ("time" < sqlc.arg('time')::time + sqlc.arg('duration')::interval AND sqlc.arg('time')::time + sqlc.arg('duration')::interval < "time" + "duration")
  )
LIMIT 1;

-- name: DeleteAppointment :execrows
DELETE FROM "appointments"
WHERE "id" = $1;
