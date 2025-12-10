-- name: CreateUser :one
INSERT INTO users (
    name,
    strength,
    inteligence,
    focus,
    speed,
    endurance
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id;

-- name: UpdateUserStats :one
UPDATE users
SET
    strength    = COALESCE($2, strength),
    inteligence = COALESCE($3, inteligence),
    focus       = COALESCE($4, focus),
    speed       = COALESCE($5, speed),
    endurance   = COALESCE($6, endurance)
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
