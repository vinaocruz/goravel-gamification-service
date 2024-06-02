CREATE TABLE players (
  id SERIAL PRIMARY KEY NOT NULL,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);
