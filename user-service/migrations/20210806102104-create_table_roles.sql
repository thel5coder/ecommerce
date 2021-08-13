-- +migrate Up
CREATE TABLE IF NOT EXISTS "roles" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "name" varchar(100) NOT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS "roles";