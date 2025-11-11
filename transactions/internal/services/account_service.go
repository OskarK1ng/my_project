package services

import (
	"context"
	"errors"
	"transactions/internal/repository"
)

func Deposit(ctx context.Context, userID string, amount float64) error {
	acc, err := repository.GetAccountByUserID(ctx, userID)
	if err != nil {
		return err
	}

	acc.Balance += amount
	return repository.UpdateBalance(ctx, userID, acc.Balance)
}

func Withdraw(ctx context.Context, userID string, amount float64) error {
	acc, err := repository.GetAccountByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if acc.Balance < amount {
		return errors.New("insufficient funds")
	}

	acc.Balance -= amount
	return repository.UpdateBalance(ctx, userID, acc.Balance)
}
