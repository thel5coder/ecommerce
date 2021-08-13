-- +migrate Up
CREATE TABLE IF NOT EXISTS "carts"
(
    "id"         char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "user_id"    char(36)                                NOT NULL,
    "product_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
    "price"      numeric                                 NOT NULL,
    "quantity"   int2                                    NOT NULL,
    "sub_total"  numeric                                 NOT NULL,
    "name"       varchar(255) COLLATE "pg_catalog"."default",
    "sku"        varchar(15) COLLATE "pg_catalog"."default",
    "category"   varchar(100),
    "created_at" timestamp                               NOT NULL,
    "updated_at" timestamp                               NOT NULL,
    "deleted_at" timestamp
);


-- +migrate Down
DROP TABLE IF EXISTS "carts";