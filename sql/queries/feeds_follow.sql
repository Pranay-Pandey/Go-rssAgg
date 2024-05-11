-- name: CreateFollow :one
INSERT INTO feeds_follow (id, created_at, updated_at, feed_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFollows :many
SELECT * from feeds_follow WHERE user_id = $1;