-- +migrate Up
CREATE TABLE IF NOT EXISTS "products"
(
    "id"             char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "category_id"    char(36)     NOT NULL,
    "name"           varchar(255) NOT NULL,
    "sku"            varchar(15)  NOT NULL,
    "price"          numeric      NOT NULL,
    "discount"       numeric,
    "main_image_key" varchar(255),
    "created_at"     timestamp    NOT NULL,
    "updated_at"     timestamp    NOT NULL,
    "deleted_at"     timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "products";