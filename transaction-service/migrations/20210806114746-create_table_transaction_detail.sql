
-- +migrate Up
CREATE TABLE IF NOT EXISTS "transaction_details" (
                                       "id" char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
                                       "transaction_id" char(36) NOT NULL,
                                       "name" varchar(255) NOT NULL,
                                       "sku" varchar(15) NOT NULL,
                                       "category" json NOT NULL,
                                       "price" numeric,
                                       "discount" numeric,
                                       "quantity" int2,
                                       "sub_total" numeric,
                                       "created_at" timestamp NOT NULL,
                                       "updated_at" timestamp NOT NULL,
                                       "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "transaction_details";