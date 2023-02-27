-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS subreddits (
  id        UUID DEFAULT  uuid_generate_v4(),
  name      VARCHAR(255)  NOT NULL UNIQUE,

  created_at TIMESTAMP DEFAULT NOW(),

  PRIMARY KEY (id)
);

INSERT INTO subreddits(name) 
VALUES
('lifeofnorman'),
('talesfromtechsupport'),
('talesfromretail'),
('glitch_in_the_matrix'),
('no_sleep'),
('UnresolvedMysteries'),
('shortscarystories'),
('psycho_alpaca'),
('relationships'), 
('AmItheAsshole') 
ON CONFLICT DO NOTHING
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subreddits;
-- +goose StatementEnd
