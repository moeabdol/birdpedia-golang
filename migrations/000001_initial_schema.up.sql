CREATE TABLE birds (
  id bigserial PRIMARY KEY,
  species varchar NOT NULL,
  description text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now()),
  updated_at timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON birds (species);
