-- name: UpdateUser :exec
UPDATE users SET 
name=$1,
username=$2,
email=$3,
mobile=$4 
WHERE id = $5;