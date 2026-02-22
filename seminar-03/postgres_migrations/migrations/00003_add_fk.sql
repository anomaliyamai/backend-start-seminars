-- +goose Up
ALTER TABLE posts
ADD CONSTRAINT posts_user_fk
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE;

-- +goose Down
ALTER TABLE posts
DROP CONSTRAINT posts_user_fk;
