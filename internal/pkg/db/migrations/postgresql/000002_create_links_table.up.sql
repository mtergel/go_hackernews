CREATE TABLE IF NOT EXISTS links(
  id serial PRIMARY KEY,
  title VARCHAR (255),
  address VARCHAR (255),
  user_id INTEGER NOT NULL,
  CONSTRAINT user_fk FOREIGN KEY(user_id) REFERENCES users(id)
)
