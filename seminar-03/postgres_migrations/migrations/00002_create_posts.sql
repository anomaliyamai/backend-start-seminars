-- +goose Up
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title TEXT NOT NULL
);

-- +goose Down
DROP TABLE posts;
