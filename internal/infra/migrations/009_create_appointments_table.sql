CREATE TABLE IF NOT EXISTS appointments (
    "id"                UUID   PRIMARY KEY        NOT NULL   DEFAULT gen_random_uuid(),
    "customer_id"       UUID                      NOT NULL,
    "specialist_id"     UUID                      NOT NULL,
    "service_name_id"   UUID                      NOT NULL,
    "price"             INT                       NOT NULL, -- in cents
    "duration"          INTERVAL                  NOT NULL, -- in minutes
    "date"              DATE                      NOT NULL,
    "time"              TIME                      NOT NULL,
    "status"            INT                       NOT NULL, -- NONE, PENDING, REALIZED, CANCELED
    "notified_at"       TIMESTAMP WITH TIME ZONE      NULL   DEFAULT NULL,
    "notified_by"       UUID                          NULL   DEFAULT NULL,

    FOREIGN KEY ("customer_id")      REFERENCES customers("id"),
    FOREIGN KEY ("specialist_id")    REFERENCES specialists("id"),
    FOREIGN KEY ("service_name_id")  REFERENCES service_names("id")
);

-- sorted down
CREATE INDEX appointments_date_idx ON appointments("date" DESC);

---- create above / drop below ----

DROP INDEX IF EXISTS appointments_date_idx;

DROP TABLE IF EXISTS appointments;
