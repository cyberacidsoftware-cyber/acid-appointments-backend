-- SELECT "sn"."id", "sn"."name", 
--         "sn"."specialization_id", "sz"."name" as "specialization_name"
-- FROM "service_names" "sn"
-- JOIN "specializations" "sz" ON "sn"."specialization_id" = "sz"."id"
-- WHERE true
--    AND ($1::text = '' OR LOWER(unaccent("sz"."name")) LIKE '%' || LOWER(unaccent($1)) || '%')
--    AND ($2::text = '' OR LOWER(unaccent("sn"."name")) LIKE '%' || LOWER(unaccent($2)) || '%');

-- name: GetServiceAvailableByID :one
SELECT "sn"."id" as "service_name_id", "sn"."name" as "service_name",
    "sp"."id" as "specialization_id", "sp"."name" as "specialization"
FROM "service_names" "sn"
JOIN "specializations" "sp" ON "sn"."specialization_id" = "sp"."id"
WHERE "sn"."id" = $1;
