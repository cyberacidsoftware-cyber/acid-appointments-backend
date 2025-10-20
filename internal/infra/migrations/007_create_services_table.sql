CREATE TABLE IF NOT EXISTS services (
    "id"                        UUID      PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "service_name_id" UUID                   NOT NULL,
    "specialist_id"             UUID                   NOT NULL,
    "price"                     INT                    NOT NULL, -- in cents
    "duration"                  INTERVAL               NOT NULL, -- in minutes

    FOREIGN KEY ("service_name_id")  REFERENCES service_names("id") ON DELETE CASCADE,
    FOREIGN KEY ("specialist_id")    REFERENCES specialists("id")   ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS services_service_unique_idx  ON services("service_name_id", "specialist_id");
CREATE        INDEX IF NOT EXISTS services_specialist_id_idx   ON services("specialist_id");
CREATE        INDEX IF NOT EXISTS services_service_name_id_idx ON services("service_name_id");

---- create above / drop below ----

DROP INDEX IF EXISTS services_service_unique_idx;
DROP INDEX IF EXISTS services_specialist_id_idx;
DROP INDEX IF EXISTS services_service_name_id_idx;

DROP TABLE IF EXISTS services;
