CREATE TABLE IF NOT EXISTS leads(
    id serial PRIMARY KEY,
    first_name VARCHAR (128) NOT NULL,
    last_name VARCHAR (128) NOT NULL,
    email VARCHAR (128) UNIQUE NOT NULL,
    gender text,
    timezone varchar(255)
);