package services

import (
	"context"
	"errors"
	"transactions/internal/repository"
)

func Deposit(ctx context.Context, accountID int64, amount float64) error {
	acc, err := repository.GetAccountByID(ctx, accountID)
	if err != nil {
		return err
	}

	acc.Balance += amount

	return repository.UpdateBalance(ctx, acc.ID, acc.Balance)
}

func Withdraw(ctx context.Context, accountID int64, amount float64) error {
	acc, err := repository.GetAccountByID(ctx, accountID)
	if err != nil {
		return err
	}

	if acc.Balance < amount {
		return errors.New("insufficient funds")
	}

	acc.Balance -= amount

	return repository.UpdateBalance(ctx, acc.ID, acc.Balance)
}
