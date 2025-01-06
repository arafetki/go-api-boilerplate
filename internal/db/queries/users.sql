-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO users (email,name) VALUES ($1,$2);