CREATE TABLE IF NOT EXISTS songs (
   set_id uuid REFERENCES sets(id),
   id uuid NOT NULL PRIMARY KEY,
   name VARCHAR(60)
);