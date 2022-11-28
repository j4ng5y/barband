ALTER TABLE songs
    ADD COLUMN IF NOT EXISTS title VARCHAR(60),
    ADD COLUMN IF NOT EXISTS genre VARCHAR(60),
    ADD COLUMN IF NOT EXISTS subgenre VARCHAR(60),
    ADD COLUMN IF NOT EXISTS primary_artist VARCHAR(60),
    ADD COLUMN IF NOT EXISTS featured_artists VARCHAR(60),
    ADD COLUMN IF NOT EXISTS composer VARCHAR(60),
    ADD COLUMN IF NOT EXISTS publisher VARCHAR(60),
    ADD COLUMN IF NOT EXISTS producers VARCHAR(60),
    ADD COLUMN IF NOT EXISTS additional_contributors VARCHAR(60),
    ADD COLUMN IF NOT EXISTS explicit_content BOOLEAN,
    ADD COLUMN IF NOT EXISTS lyrics_language VARCHAR(60),
    ADD COLUMN IF NOT EXISTS lyrics_publisher VARCHAR(60),
    ADD COLUMN IF NOT EXISTS composition_owner VARCHAR(60),
    ADD COLUMN IF NOT EXISTS year_of_composition INTEGER,
    ADD COLUMN IF NOT EXISTS master_recording_owner VARCHAR(60),
    ADD COLUMN IF NOT EXISTS year_of_recording INTEGER,
    ADD COLUMN IF NOT EXISTS release_language VARCHAR(60);
UPDATE songs SET title = name;
