// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: save_refresh_token.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const saveRefreshToken = `-- name: SaveRefreshToken :one
INSERT INTO refresh_token(token, created_at, updated_at, user_id, expires_at)
VALUES ($1, NOW(), NOW(), $2, $3)
RETURNING token, created_at, updated_at, user_id, expires_at, revoked_at
`

type SaveRefreshTokenParams struct {
	Token     string
	UserID    uuid.UUID
	ExpiresAt time.Time
}

func (q *Queries) SaveRefreshToken(ctx context.Context, arg SaveRefreshTokenParams) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, saveRefreshToken, arg.Token, arg.UserID, arg.ExpiresAt)
	var i RefreshToken
	err := row.Scan(
		&i.Token,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.ExpiresAt,
		&i.RevokedAt,
	)
	return i, err
}