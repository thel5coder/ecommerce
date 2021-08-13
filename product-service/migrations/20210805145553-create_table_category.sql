-- +migrate Up
CREATE TABLE IF NOT EXISTS "categories"
(
    "id"         char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "name"       varchar(60)  NOT NULL,
    "slug"       varchar(255) NOT NULL,
    "created_at" timestamp    NOT NULL,
    "updated_at" timestamp    NOT NULL,
    "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "categories";