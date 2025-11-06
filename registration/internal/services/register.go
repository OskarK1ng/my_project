package services

import (
	"context"
	"fmt"
	"registration/internal/db"
	"registration/internal/models"
	"time"

	"github.com/google/uuid"
)

func RegisterUser(ctx context.Context, user *models.User) error {
	user.UserID = uuid.New().String()
	user.CreatedAt = time.Now()

	query := `
	INSERT INTO users (
		user_id,
		first_name,
		last_name,
		email,
		phone,
		created_at
		)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	)
	`

	_, err := db.DB.Exec(ctx, query,
		user.UserID, user.FirstName, user.LastName, user.Email, user.Phone, user.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}
