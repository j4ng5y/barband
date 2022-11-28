CREATE TABLE IF NOT EXISTS band_members (
  band_id uuid REFERENCES bands(id),
  id uuid NOT NULL PRIMARY KEY,
  prefix VARCHAR(12),
  first_name VARCHAR(60) NOT NULL,
  middle_name VARCHAR(60),
  last_name VARCHAR(60),
  suffix VARCHAR(12),
  nick_name VARCHAR(60),
  phone_number INTEGER
);