-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `todos`(
    `id` INTEGER PRIMARY KEY, 
    `user_id` INTEGER NOT NULL, 
    'name' TEXT NOT NULL,
    'description' TEXT NULL,
    FOREIGN KEY (`user_id`) 
      REFERENCES `users` (`id`) 
         ON DELETE CASCADE 
         ON UPDATE NO ACTION
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `todos`;
-- +goose StatementEnd
