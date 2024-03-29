-- name: GetProfiles :many
SELECT region, c.access_id, c.secret_key FROM profiles p 
INNER JOIN credentials c ON c.id = p.cred_id 
WHERE p.username = $1 LIMIT $2;

-- name: CreateProfile :one
INSERT INTO profiles (id, description, region, cred_id, username)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;