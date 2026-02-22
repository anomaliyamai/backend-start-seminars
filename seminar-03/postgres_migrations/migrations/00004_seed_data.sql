-- +goose Up
INSERT INTO users (name, email) VALUES
('Alice', 'alice@mail.com'),
('Bob', 'bob@mail.com');

INSERT INTO posts (user_id, title) VALUES
(1, 'Hello World'),
(1, 'Go + Postgres'),
(2, 'Bob first post');

-- +goose Down
DELETE FROM posts;
DELETE FROM users;
