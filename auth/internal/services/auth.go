package services

import (
	"auth/internal/db"
	"auth/internal/models"
	"auth/internal/security"
	"context"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (string, error) {
	var user models.User

	// Ищем пользователя в БД
	query := `
		SELECT
			user_id,
			user_email,
			user_password_hash
		FROM users
		WHERE user_email = $1
	`

	err := db.DB.QueryRow(context.Background(), query, email).Scan(
		&user.UserID,
		&user.UserEmail,
		&user.UserPasswordHash,
	)
	if err != nil {
		log.Printf("user not found: %v", err)

		return "", errors.New("invalid credentials")
	}

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(user.UserPasswordHash), []byte(password)); err != nil {
		log.Printf("invalid password for user %s", email)

		return "", errors.New("invalid credentials")
	}

	// Генерируем JWT
	token, err := security.GenerateJWT(user.UserEmail)
	if err != nil {
		log.Printf("error generating token: %v", err)

		return "", err
	}

	return token, nil
}
