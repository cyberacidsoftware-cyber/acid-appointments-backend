CREATE TABLE IF NOT EXISTS specialist_hours (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "specialist_id"     UUID                      NOT NULL,
    "weekday"           INT                       NOT NULL, -- 0 = Sunday, 1 = Monday, ..., 6 = Saturday
    "start_time"        TIME                      NOT NULL,
    "end_time"          TIME                      NOT NULL,
    
    FOREIGN KEY ("specialist_id") REFERENCES specialists("id") ON DELETE CASCADE
);

---- create above / drop below ----

DROP TABLE IF EXISTS specialist_hours;