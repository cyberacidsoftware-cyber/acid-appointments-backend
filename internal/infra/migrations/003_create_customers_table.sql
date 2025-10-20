CREATE TABLE IF NOT EXISTS customers (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "email"             VARCHAR(255)                  NULL,
    "phone"             VARCHAR(255)              NOT NULL, -- UNIQUE
    "birthdate"         DATE                      NOT NULL,
    "cpf"               VARCHAR(11)               NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_customers_phone" ON customers("phone");

---- create above / drop below ----

DROP INDEX IF EXISTS idx_customers_phone;

DROP TABLE IF EXISTS customers;
