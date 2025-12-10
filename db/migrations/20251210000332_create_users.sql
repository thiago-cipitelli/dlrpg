-- +goose Up
CREATE TABLE users (
  id            BIGSERIAL   PRIMARY KEY,
  name          text        NOT NULL,
  strength      integer     DEFAULT 0,
  inteligence   integer     DEFAULT 0,
  focus         integer     DEFAULT 0,
  speed         integer     DEFAULT 0,
  endurance     integer     DEFAULT 0
);

-- +goose Down
DROP TABLE users;
