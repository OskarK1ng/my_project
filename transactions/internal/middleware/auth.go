package middleware

import (
	"net/http"
	"strings"
	"transactions/internal/config"
	"transactions/internal/security"
)

// AuthMiddleware проверяет JWT токен
func AuthMiddleware(cfg *config.Config, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Ожидается: Authorization: Bearer <token>
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		token := parts[1]

		_, err := security.ValidateJWT(token, cfg.JWTSecret)
		if err != nil {
			http.Error(w, "invalid or expired token", http.StatusUnauthorized)
			return
		}

		// если токен валиден — продолжаем
		next(w, r)
	}
}
