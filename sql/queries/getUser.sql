-- name: GetUserAuth :one
SELECT * FROM users WHERE username = $1 AND password = $2;
