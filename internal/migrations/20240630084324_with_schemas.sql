-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA user_schema;
CREATE SCHEMA payment_schema;
CREATE SCHEMA product_schema;

ALTER TABLE public.users SET SCHEMA user_schema;
ALTER TABLE public.balances SET SCHEMA payment_schema;
ALTER TABLE public.operations SET SCHEMA payment_schema;
ALTER TABLE public.cards SET SCHEMA product_schema;
ALTER TABLE public.deposits SET SCHEMA product_schema;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_schema.users SET SCHEMA public;
ALTER TABLE payment_schema.balances SET SCHEMA public;
ALTER TABLE payment_schema.operations SET SCHEMA public;
ALTER TABLE product_schema.cards SET SCHEMA public;
ALTER TABLE product_schema.deposits SET SCHEMA public;

DROP SCHEMA IF EXISTS user_schema;
DROP SCHEMA IF EXISTS payment_schema;
DROP SCHEMA IF EXISTS product_schema;
-- +goose StatementEnd
