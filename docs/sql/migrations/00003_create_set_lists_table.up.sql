CREATE TABLE IF NOT EXISTS set_lists (
   band_id uuid REFERENCES bands(id),
   id uuid NOT NULL PRIMARY KEY,
   name VARCHAR(60)
);