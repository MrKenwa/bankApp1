-- +goose Up
-- +goose StatementBegin
ALTER TABLE product_schema.cards
    ADD CONSTRAINT user_fk1 FOREIGN KEY (user_id) REFERENCES user_schema.users(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE product_schema.cards
    DROP CONSTRAINT user_fk1;
-- +goose StatementEnd
