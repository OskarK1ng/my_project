package handlers

import (
	"auth/internal/config"
	"auth/internal/services"
	"context"
	"encoding/json"
	"net/http"
)

func LoginHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		if credentials.Email == "" || credentials.Password == "" {
			http.Error(w, "email and password required", http.StatusBadRequest)
			return
		}

		// Передаём cfg и context
		token, err := services.Login(context.Background(), credentials.Email, credentials.Password, cfg)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
	}
}
