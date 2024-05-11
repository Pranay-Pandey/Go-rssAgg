-- +goose Up
ALTER TABLE app_user ADD column api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea), 'hex')
);

-- +goose Down
ALTER TABLE app_user DROP COLUMN api_key;