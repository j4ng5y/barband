CREATE TABLE IF NOT EXISTS lyrics (
  song_id uuid REFERENCES songs(id),
  id uuid PRIMARY KEY NOT NULL,
  part VARCHAR(60),
  lyrics TEXT
);

CREATE TABLE IF NOT EXISTS tabs (
  song_id uuid REFERENCES songs(id),
  id uuid PRIMARY KEY NOT NULL,
  part VARCHAR(60),
  tabs TEXT
);