package services

import (
	"context"
	"errors"
	"fmt"
	"transactions/internal/repository"
)

func GetBalance(ctx context.Context, userID string) (float64, error) {
	if userID == "" {
		return 0, fmt.Errorf("user id is required")
	}

	acc, err := repository.GetBalance(ctx, userID)
	if err != nil {
		return 0, err
	}

	return acc.Balance, nil
}

func Deposit(ctx context.Context, userID string, amount float64) error {
	if userID == "" {
		return fmt.Errorf("user id is required")
	}

	acc, err := repository.GetBalance(ctx, userID)
	if err != nil {
		return err
	}

	acc.Balance += amount

	return repository.UpdateBalance(ctx, userID, acc.Balance)
}

func Withdraw(ctx context.Context, userID string, amount float64) error {
	if userID == "" {
		return fmt.Errorf("user id is required")
	}

	acc, err := repository.GetBalance(ctx, userID)
	if err != nil {
		return err
	}

	if acc.Balance < amount {
		return errors.New("insufficient funds")
	}

	acc.Balance -= amount

	return repository.UpdateBalance(ctx, userID, acc.Balance)
}
