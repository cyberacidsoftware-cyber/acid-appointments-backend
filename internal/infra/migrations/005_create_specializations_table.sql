CREATE TABLE IF NOT EXISTS specializations (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL  -- UNIQUE
);

CREATE UNIQUE INDEX IF NOT EXISTS specializations_name_idx ON specializations("name");

---- create above / drop below ----

DROP INDEX IF EXISTS specializations_name_idx;

DROP TABLE IF EXISTS specializations;
