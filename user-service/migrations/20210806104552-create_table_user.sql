-- +migrate Up
CREATE TABLE IF NOT EXISTS "users"
(
    "id"         char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "email"      varchar(100) NOT NULL,
    "first_name" varchar(255) NOT NULL,
    "last_name"  varchar(255),
    "password"   varchar(128) NOT NULL,
    "role_id"    int          NOT NULL,
    "created_at" timestamp    NOT NULL,
    "updated_at" timestamp    NOT NULL,
    "deleted_at" timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "users";