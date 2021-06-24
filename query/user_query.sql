-- WORK
-- name: CreateUser :one
INSERT INTO users (username, password)
VALUES ($1, $2) 
RETURNING username, password;

-- WORK
-- name: UpdateUserToMerchant :one
UPDATE users SET role = 'merchant'
WHERE username = $1
RETURNING id;

-- WORK
-- name: GetUsers :many
SELECT * FROM users;

-- WORK
-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- WORK
-- name: UpdateUserPassword :exec
UPDATE users SET password = $2
WHERE id = $1;

-- name: GetUserPassword :one
SELECT password FROM users
WHERE username = $1;

-- WORK
-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;

-- name: IsUsernameExist :one
SELECT id FROM users
WHERE username = $1;