-- name: FindUserByEmail :one
SELECT id, email, hashed_password, role
FROM users
WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users (id, email, hashed_password, role)
VALUES ($1, $2, $3, $4);

-- name: ExistsUserByEmail :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);
