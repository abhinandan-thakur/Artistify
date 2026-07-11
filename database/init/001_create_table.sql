CREATE TABLE IF NOT EXISTS albums (
    id SERIAL PRIMARY KEY,
    album_name TEXT NOT NULL,
    artist TEXT NOT NULL,
    sales BIGINT NOT NULL,
    rating NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    type TEXT NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW()
);