-- name: GetChirpsByAuthorDesc :many
SELECT * FROM chirps
WHERE user_id = $1
ORDER BY created_at DESC;