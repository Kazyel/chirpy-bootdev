-- +goose Up
CREATE TABLE
    refresh_token (
        token TEXT PRIMARY KEY NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        user_id UUID NOT NULL,
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
        expires_at TIMESTAMP NOT NULL,
        revoked_at TIMESTAMP
    );

-- +goose Down
DROP TABLE refresh_token;