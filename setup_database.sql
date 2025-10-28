-- Create extension first
CREATE EXTENSION IF NOT EXISTS unaccent;

-- 1. Create admins table
CREATE TABLE IF NOT EXISTS "admins" (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "email"             VARCHAR(255)              NOT NULL,
    "password"          VARCHAR(255)              NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_admins_email" ON admins("email");

-- 2. Create secretaries table
CREATE TABLE IF NOT EXISTS secretaries (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "email"             VARCHAR(255)              NOT NULL,
    "password"          VARCHAR(255)              NOT NULL,
    "phone"             VARCHAR(255)              NOT NULL,
    "birthdate"         DATE                      NULL,
    "cpf"               VARCHAR(11)               NOT NULL,
    "cnpj"              VARCHAR(14)               NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS secretaries_email_idx ON secretaries("email");

-- 3. Create customers table
CREATE TABLE IF NOT EXISTS customers (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "email"             VARCHAR(255)              NULL,
    "phone"             VARCHAR(255)              NOT NULL,
    "birthdate"         DATE                      NULL,
    "cpf"               VARCHAR(11)               NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS customers_cpf_idx ON customers("cpf");

-- 4. Create specialists table
CREATE TABLE IF NOT EXISTS specialists (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "email"             VARCHAR(255)              NOT NULL,
    "phone"             VARCHAR(255)              NOT NULL,
    "birthdate"         DATE                      NULL,
    "cpf"               VARCHAR(11)               NOT NULL,
    "cnpj"              VARCHAR(14)               NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS specialists_email_idx ON specialists("email");

-- 5. Create specializations table
CREATE TABLE IF NOT EXISTS specializations (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS specializations_name_idx ON specializations("name");

-- 6. Create service_names table
CREATE TABLE IF NOT EXISTS service_names (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "name"              VARCHAR(255)              NOT NULL,
    "specialization_id" UUID                      NOT NULL,
    FOREIGN KEY ("specialization_id") REFERENCES specializations("id") ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS service_names_name_idx ON service_names("name");
CREATE INDEX IF NOT EXISTS service_names_specialization_id_idx ON service_names("specialization_id");

-- 7. Create services table
CREATE TABLE IF NOT EXISTS services (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "service_name_id"   UUID                      NOT NULL,
    "specialist_id"     UUID                      NOT NULL,
    "price"             INTEGER                   NOT NULL,
    "duration"          INTERVAL                  NOT NULL,
    FOREIGN KEY ("service_name_id") REFERENCES service_names("id") ON DELETE CASCADE,
    FOREIGN KEY ("specialist_id") REFERENCES specialists("id") ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS services_service_unique_idx ON services("service_name_id", "specialist_id");
CREATE INDEX IF NOT EXISTS services_specialist_id_idx ON services("specialist_id");
CREATE INDEX IF NOT EXISTS services_service_name_id_idx ON services("service_name_id");

-- 8. Create specialist_hours table
CREATE TABLE IF NOT EXISTS specialist_hours (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "specialist_id"     UUID                      NOT NULL,
    "weekday"           INTEGER                   NOT NULL,
    "start_time"        TIME                      NOT NULL,
    "end_time"          TIME                      NOT NULL,
    FOREIGN KEY ("specialist_id") REFERENCES specialists("id") ON DELETE CASCADE
);

-- 9. Create appointments table
CREATE TABLE IF NOT EXISTS appointments (
    "id"                UUID         PRIMARY KEY  NOT NULL   DEFAULT gen_random_uuid(),
    "customer_id"       UUID                      NOT NULL,
    "specialist_id"     UUID                      NOT NULL,
    "service_name_id"   UUID                      NOT NULL,
    "price"             INTEGER                   NOT NULL,
    "duration"          INTERVAL                  NOT NULL,
    "date"              DATE                      NOT NULL,
    "time"              TIME                      NOT NULL,
    "status"            INTEGER                   NOT NULL   DEFAULT 0,
    "notified_at"       TIMESTAMPTZ               NULL,
    "notified_by"       UUID                      NULL,
    FOREIGN KEY ("customer_id") REFERENCES customers("id") ON DELETE CASCADE,
    FOREIGN KEY ("specialist_id") REFERENCES specialists("id") ON DELETE CASCADE,
    FOREIGN KEY ("service_name_id") REFERENCES service_names("id") ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS appointments_date_idx ON appointments("date");