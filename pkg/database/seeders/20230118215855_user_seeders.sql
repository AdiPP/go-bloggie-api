-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, username, password, email)
VALUES (1, 'John', '12345', 'john@mail.com');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
