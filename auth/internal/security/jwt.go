package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email                string `json:"email"`
	jwt.RegisteredClaims        // встраивание структур
}

// ttl - time to life
func GenerateJWT(email string, secret string, ttlMinutes int) (string, error) {
	// Установим TTL (5 min standart)
	expTime := time.Now().Add(time.Duration(ttlMinutes) * time.Minute)

	// Создадим claims
	claims := Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// Создаем JWT-token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	// Подписываем token (генерируем полноценный token с checksum)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
