-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id integer PRIMARY KEY,
    username varchar(50) NOT NULL,
    password varchar(255) NOT NULL,
    email varchar(50) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
