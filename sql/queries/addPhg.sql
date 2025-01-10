-- name: AddPhotographer :exec
INSERT INTO photographers (phgname,phgusername,phgmobile,phgpassword)
VALUES ($1,$2,$3,$4)
RETURNING *;
