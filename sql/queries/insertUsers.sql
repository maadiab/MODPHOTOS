-- name: CreateUser :exec
INSERT INTO users (name,username,email,mobile,password)
VALUES($1,$2,$3,$4,$5)
RETURNING *;
