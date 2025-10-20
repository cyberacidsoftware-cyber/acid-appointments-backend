CREATE TABLE IF NOT EXISTS "admins" (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "email"             VARCHAR(255)              NOT NULL, -- UNIQUE
    "password"          VARCHAR(255)              NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_admins_email" ON admins("email");

---- create above / drop below ----

DROP INDEX IF EXISTS idx_admins_email;

DROP TABLE IF EXISTS "admins";
