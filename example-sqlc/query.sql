
-- name: GetUnderAge :many
SELECT * FROM players WHERE age <= ?;

-- name: AddAge :exec
UPDATE players SET age = age + ?;
