-- name: CreateCredential :one
INSERT INTO credentials (id, access_id, secret_key)
VALUES ($1, $2, $3)
returning *;