-- name: GetUserAuth :one
SELECT * FROM users WHERE username = $1;
