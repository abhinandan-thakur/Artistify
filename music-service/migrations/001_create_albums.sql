CREATE TABLE albums (
    id SERIAL PRIMARY KEY,
    album_name TEXT NOT NULL,
    artist TEXT NOT NULL,
    sales BIGINT NOT NULL,
    rating NUMERIC(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);