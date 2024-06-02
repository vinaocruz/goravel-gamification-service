CREATE TABLE player_events (
  id SERIAL PRIMARY KEY NOT NULL,
  player_id INT NOT NULL,
  event_id INT NOT NULL,
  points INT NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  deleted_at timestamp,
  FOREIGN KEY (player_id) REFERENCES players(id),
  FOREIGN KEY (event_id) REFERENCES events(id)
);
