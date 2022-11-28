CREATE TABLE IF NOT EXISTS sets (
  set_list_id uuid REFERENCES set_lists(id),
  id uuid NOT NULL PRIMARY KEY,
  name VARCHAR(60)
);