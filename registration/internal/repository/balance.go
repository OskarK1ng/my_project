package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// CreateBalance — создаёт запись с нулевым балансом для нового пользователя.
func CreateBalance(ctx context.Context, tx pgx.Tx, userID string) error {
	query := `
		INSERT INTO balance (user_id, balance)
		VALUES ($1, 0)
	`
	_, err := tx.Exec(ctx, query, userID)
	return err
}
