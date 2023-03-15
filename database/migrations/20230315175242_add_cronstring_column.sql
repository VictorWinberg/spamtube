-- +goose Up
-- +goose StatementBegin
alter TABLE subreddits  ADD column if not exists cron_string VARCHAR(255) DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter TABLE subreddits DROP column if exists cron_string VARCHAR(255);
-- +goose StatementEnd
