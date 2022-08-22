-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE posts (
  id         UUID DEFAULT  uuid_generate_v4(),
  source_id  VARCHAR(255)  NOT NULL UNIQUE,
  dest_id    VARCHAR(255)  NOT NULL UNIQUE,

  created_at TIMESTAMP DEFAULT NOW(),

  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts CASCADE;
-- +goose StatementEnd
