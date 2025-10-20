CREATE TABLE IF NOT EXISTS service_names (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "specialization_id" UUID                      NOT NULL,

    FOREIGN KEY ("specialization_id") REFERENCES specializations("id") ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS service_names_name_idx              ON service_names("name");
CREATE        INDEX IF NOT EXISTS service_names_specialization_id_idx ON service_names("specialization_id");

---- create above / drop below ----

DROP INDEX IF EXISTS service_names_name_idx;
DROP INDEX IF EXISTS service_names_specialization_id_idx;

DROP TABLE IF EXISTS service_names;
