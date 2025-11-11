package repository

import (
	"context"
	"fmt"
	"transactions/internal/db"
	"transactions/internal/models"
)

func CreateAccount(ctx context.Context, userID string) (models.Account, error) {
	query := `
		INSERT INTO transactions (user_id, balance, created_at)
		VALUES ($1, 0, NOW())
		RETURNING id, user_id, balance, created_at
	`

	var acc models.Account
	err := db.DB.QueryRow(ctx, query, userID).Scan(
		&acc.ID,
		&acc.UserID,
		&acc.Balance,
		&acc.CreatedAt,
	)
	return acc, err
}

func GetAccountByUserID(ctx context.Context, userID string) (models.Account, error) {
	var acc models.Account
	query := `
        SELECT id, user_id, balance, created_at
        FROM transactions
        WHERE user_id=$1
    `
	err := db.DB.QueryRow(ctx, query, userID).Scan(
		&acc.ID,
		&acc.UserID,
		&acc.Balance,
		&acc.CreatedAt,
	)
	return acc, err
}

func UpdateBalance(ctx context.Context, userID string, newBalance float64) error {
	query := `
        UPDATE transactions
        SET balance=$1
        WHERE user_id=$2
    `
	cmd, err := db.DB.Exec(ctx, query, newBalance, userID)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("account not found")
	}
	return nil
}
