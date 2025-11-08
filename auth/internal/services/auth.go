package services

import (
	"auth/internal/config"
	"auth/internal/db"
	"auth/internal/security"
	"context"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Login проверяет почту и пароль, затем генерирует JWT при успехе
func Login(ctx context.Context, email, password string, cfg *config.Config) (string, error) {
	// Проверяем, что поля не пустые
	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	var storedHash string
	var storedEmail string

	// Ищем пользователя по email
	query := `
		SELECT user_email, user_password_hash 
		FROM users 
		WHERE user_email = $1
	`
	err := db.DB.QueryRow(ctx, query, email).Scan(&storedEmail, &storedHash)
	if err != nil {
		log.Printf("User not found or query failed: %v", err)
		return "", errors.New("invalid email or password")
	}

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)); err != nil {
		log.Printf("Password mismatch for user %s", email)
		return "", errors.New("invalid email or password")
	}

	// Генерация JWT
	token, err := security.GenerateJWT(email, cfg.JWTSecret, cfg.JWTTTLMinutes)
	if err != nil {
		return "", fmt.Errorf("failed to generate jwt: %w", err)
	}

	log.Printf("User %s logged in successfully", email)
	return token, nil
}
