DROP TABLE IF EXISTS users;
CREATE TABLE  users (
  id   SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  age  INTEGER NOT NULL
);
