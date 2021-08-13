
-- +migrate Up

CREATE TYPE "transaction_status_enum" AS ENUM (
    'cart',
    'on_going',
    'success',
    'canceled'
    );

-- +migrate Down
