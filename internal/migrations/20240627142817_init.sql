-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    lastname VARCHAR(40) NOT NULL,
    patronymic VARCHAR(40),
    email VARCHAR(50) NOT NULL,
    password VARCHAR(256) NOT NULL,
    passport_number VARCHAR(256) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS cards(
    card_id BIGSERIAL PRIMARY KEY,
    card_number BIGINT NOT NULL UNIQUE,
    user_id BIGINT NOT NULL,
    card_type VARCHAR(64) NOT NULL DEFAULT 'debit card',
    pin_code TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS balances(
    balance_id BIGSERIAL PRIMARY KEY,
    card_id BIGINT,
    deposit_id BIGINT,
    amount BIGINT CHECK ( amount >= 0 ) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT card_fk FOREIGN KEY (card_id) REFERENCES cards(card_id),
    CONSTRAINT deposit_fk FOREIGN KEY (deposit_id) REFERENCES deposits(deposit_id)
);

CREATE TABLE IF NOT EXISTS operations(
    operation_id BIGSERIAL PRIMARY KEY,
    sender_balance_id BIGINT,
    receiver_balance_id BIGINT,
    amount BIGINT NOT NULL,
    operation_type VARCHAR(32) NOT NULL DEFAULT 'transfer',
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT sender_fk FOREIGN KEY(sender_balance_id) REFERENCES balances(balance_id),
    CONSTRAINT receiver_fk FOREIGN KEY(receiver_balance_id) REFERENCES balances(balance_id)
);

CREATE TABLE IF NOT EXISTS deposits(
    deposit_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    deposit_type VARCHAR(32) NOT NULL DEFAULT 'deposit',
    interest_rate NUMERIC NOT NULL,
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT user_fk FOREIGN KEY(user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS operations;
DROP TABLE IF EXISTS deposits;
DROP TABLE IF EXISTS balances;
-- +goose StatementEnd
