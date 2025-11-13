package repository

import (
	"context"
	"fmt"
	"transactions/internal/db"
	"transactions/internal/models"
)

func GetBalance(ctx context.Context, userID string) (models.Account, error) {
	var acc models.Account
	query := `
        SELECT user_id, balance
        FROM balance
        WHERE user_id=$1
    `

	err := db.DB.QueryRow(ctx, query, userID).Scan(
		&acc.UserID,
		&acc.Balance,
	)
	return acc, err
}

func UpdateBalance(ctx context.Context, userID string, newBalance float64) error {
	query := `
        UPDATE balance
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
