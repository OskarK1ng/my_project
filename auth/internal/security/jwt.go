package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Должно быть в env
const SecretJWT = "secret_key" // in real life, example: b2dg8w88)*(^&^TV)

type Claims struct {
	Email                string `json:"email"`
	jwt.RegisteredClaims        // встраивание структур
}

// ttl - time to life
func GenerateJWT(email string) (string, error) {
	// Установим TTL (5 min standart)
	expTime := time.Now().Add(5 * time.Minute)

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
	tokenString, err := token.SignedString([]byte(SecretJWT))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
