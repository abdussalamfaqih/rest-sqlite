-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `users`(
    `id` INTEGER PRIMARY KEY, 
    'name' TEXT NOT NULL,
    'password' TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `users`;
-- +goose StatementEnd
