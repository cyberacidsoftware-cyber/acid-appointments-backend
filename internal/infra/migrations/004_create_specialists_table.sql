CREATE TABLE IF NOT EXISTS specialists (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "email"             VARCHAR(255)              NOT NULL, -- UNIQUE
    "phone"             VARCHAR(255)              NOT NULL,
    "birthdate"         DATE                      NOT NULL,
    "cpf"               VARCHAR(11)               NOT NULL,
    "cnpj"              VARCHAR(14)                   NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_specialists_email" ON specialists("email");

---- create above / drop below ----

DROP INDEX IF EXISTS idx_specialists_email;

DROP TABLE IF EXISTS specialists;
