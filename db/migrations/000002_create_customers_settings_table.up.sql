CREATE TABLE IF NOT EXISTS customers_settings(
    id serial PRIMARY KEY,
    customer_id int REFERENCES customers(id) ON DELETE CASCADE NOT NULL,
    filter_wh_from timestamp,
    filter_wh_to timestamp,
    filter_priority smallint,
    filter_possibilities text[]
);