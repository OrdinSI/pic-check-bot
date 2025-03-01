-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY,
    username VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS groups (
    id BIGINT PRIMARY KEY,
    group_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS images (
    id SERIAL PRIMARY KEY,
    hash_part1 BIGINT NOT NULL,
    hash_part2 BIGINT NOT NULL,
    file_hash BYTEA NOT NULL,
    user_id BIGINT NOT NULL,
    group_id BIGINT NOT NULL,
    file_id VARCHAR(255) NOT NULL,
    message_id INT NOT NULL,
    post_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);

CREATE INDEX idx_images_hash_parts ON images (hash_part1, hash_part2);

CREATE TABLE IF NOT EXISTS reposts (
    id SERIAL PRIMARY KEY,
    image_id INT NOT NULL,
    user_id BIGINT NOT NULL,
    group_id BIGINT NOT NULL,
    repost_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS reposts;
DROP TABLE IF EXISTS images;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS groups;
-- +goose StatementEnd
