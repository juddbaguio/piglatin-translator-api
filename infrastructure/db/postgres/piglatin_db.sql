DROP TABLE IF EXISTS piglatin_requests;
CREATE TABLE IF NOT EXISTS piglatin_requests (
    id SERIAL PRIMARY KEY,
    request TEXT NOT NULL,
    translation TEXT NOT NULL
);