CREATE TABLE IF NOT EXISTS secretaries (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "email"             VARCHAR(255)              NOT NULL, -- UNIQUE
    "password"          VARCHAR(255)              NOT NULL,
    "phone"             VARCHAR(255)              NOT NULL,
    "birthdate"         DATE                      NOT NULL,
    "cpf"               VARCHAR(11)               NOT NULL,
    "cnpj"              VARCHAR(14)                   NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_secretaries_email" ON secretaries("email");

---- create above / drop below ----

DROP INDEX IF EXISTS idx_secretaries_email;

DROP TABLE IF EXISTS secretaries;
