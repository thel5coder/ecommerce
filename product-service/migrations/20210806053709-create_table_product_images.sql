-- +migrate Up
CREATE TABLE IF NOT EXISTS "product_images"
(
    "id"         char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "product_id" char(36) NOT NULL,
    "image_key"  varchar(255)
);

-- +migrate Down
DROP TABLE IF EXISTS "product_images";