package services

import (
	"context"
	"fmt"
	"registration/internal/db"
	"registration/internal/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser создает нового клиента в БД с хешированным паролем
func RegisterUser(ctx context.Context, user *models.User) error {
	user.UserID = uuid.New().String()
	user.UserCreatedAt = time.Now()

	// проверяем, что пароль не пуст
	if user.UserPassword == "" {
		return fmt.Errorf("password is required")
	}

	// Проверяем, не зарегистрирован ли уже пользователь с таким email или телефоном
	var count int
	checkQuery := `
		SELECT COUNT(*) 
		FROM users 
		WHERE user_email = $1 OR user_phone = $2
	`

	err := db.DB.QueryRow(ctx, checkQuery, user.UserEmail, user.UserPhone).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check existing user: %w", err)
	}

	if count > 0 {
		return fmt.Errorf("user with this email or phone already exists")
	}

	// хэшируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	query := `
	INSERT INTO users (
		user_id,
		user_name,
		user_last_name,
		user_email,
		user_phone,
		user_password_hash,
		user_created_at
		)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	)
	`

	_, err = db.DB.Exec(ctx, query,
		user.UserID,
		user.UserName,
		user.UserLastName,
		user.UserEmail,
		user.UserPhone,
		string(hashedPassword),
		user.UserCreatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}
