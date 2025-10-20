-- name: ListServicesEnriched :many
SELECT "s"."id", "s"."price", "s"."duration",
       "s"."specialist_id", "sp"."name" as "specialist_name",
       "s"."service_name_id", "sn"."name" as "service_name",
       "sn"."specialization_id", "sz"."name" as "specialization_name"
FROM "services" "s"
JOIN "specialists" "sp" ON "s"."specialist_id" = "sp"."id"
JOIN "service_names" "sn" ON "s"."service_name_id" = "sn"."id"
JOIN "specializations" "sz" ON "sn"."specialization_id" = "sz"."id"
WHERE true
   AND (sqlc.arg('specialist')::text = '' OR LOWER(unaccent("sp"."name")) LIKE '%' || LOWER(unaccent(sqlc.arg('specialist'))) || '%')
   AND (sqlc.arg('specialization')::text = '' OR LOWER(unaccent("sz"."name")) LIKE '%' || LOWER(unaccent(sqlc.arg('specialization'))) || '%')
   AND (sqlc.arg('service_name')::text = '' OR LOWER(unaccent("sn"."name")) LIKE '%' || LOWER(unaccent(sqlc.arg('service_name'))) || '%')
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: CountServicesEnriched :one
SELECT COUNT("s"."id")
FROM "services" "s"
JOIN "specialists" "sp" ON "s"."specialist_id" = "sp"."id"
JOIN "service_names" "sn" ON "s"."service_name_id" = "sn"."id"
JOIN "specializations" "sz" ON "sn"."specialization_id" = "sz"."id"
WHERE true
   AND (sqlc.arg('specialist')::text = '' OR LOWER(unaccent("sp"."name")) LIKE '%' || LOWER(unaccent(sqlc.arg('specialist'))) || '%')
   AND (sqlc.arg('specialization')::text = '' OR LOWER(unaccent("sz"."name")) LIKE '%' || LOWER(unaccent(sqlc.arg('specialization'))) || '%')
   AND (sqlc.arg('service_name')::text = '' OR LOWER(unaccent("sn"."name")) LIKE '%' || LOWER(unaccent(sqlc.arg('service_name'))) || '%');
