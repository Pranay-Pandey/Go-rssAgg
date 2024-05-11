-- name: CreateUser :one
INSERT INTO app_user (id, created_at, updated_at, name) 
VALUES ($1, $2, $3, $4)
RETURNING *;