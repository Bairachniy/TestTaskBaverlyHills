CREATE TABLE IF NOT EXISTS leads_settings(
    id serial PRIMARY KEY,
    lead_id integer REFERENCES leads(id) ON DELETE CASCADE NOT NULL,
    working_hours_from text,
    working_hours_to text,
    priority smallint,
    possibilities text[]
);