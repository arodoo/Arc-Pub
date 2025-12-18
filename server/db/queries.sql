-- name: FindUserByEmail :one
SELECT id, email, hashed_password, role, faction
FROM users
WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users (id, email, hashed_password, role)
VALUES ($1, $2, $3, $4);

-- name: ExistsUserByEmail :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);

-- name: GetUserProfile :one
SELECT id, email, role, faction
FROM users
WHERE id = $1;

-- name: SetUserFaction :exec
UPDATE users SET faction = $1 WHERE id = $2 AND faction IS NULL;

-- name: CreateShip :exec
INSERT INTO ships (id, user_id, ship_type, slot)
VALUES ($1, $2, $3, $4);

-- name: GetUserShips :many
SELECT id, ship_type, slot, created_at
FROM ships
WHERE user_id = $1
ORDER BY slot;

-- name: CountUserShips :one
SELECT COUNT(*) FROM ships WHERE user_id = $1;
