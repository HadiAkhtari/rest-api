CREATE TABLE IF NOT EXISTS events(
    id SERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    descripation TEXT NOT NULL,
    data TIMESTAMP NOT NULL ,
    location TEXT NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE
);