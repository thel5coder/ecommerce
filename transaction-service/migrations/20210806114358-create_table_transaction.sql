-- +migrate Up
CREATE TABLE IF NOT EXISTS "transactions"
(
    "id"                 char(36) PRIMARY KEY             DEFAULT (uuid_generate_v4()),
    "user_id"            char(36)                NOT NULL,
    "transaction_number" varchar(10)             NOT NULL,
    "status"             transaction_status_enum NOT NULL DEFAULT ('on_going'),
    "total"              numeric                 NOT NULL,
    "discount"           numeric,
    "created_at"         timestamp               NOT NULL,
    "updated_at"         timestamp               NOT NULL,
    "paid_at"            timestamp               NOT NULL,
    "canceled_at"        timestamp               NOT NULL,
    "deleted_at"         timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "transactions";