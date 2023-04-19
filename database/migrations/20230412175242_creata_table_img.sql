-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS img (
  url      VARCHAR(255)  NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS img CASCADE;
-- +goose StatementEnd