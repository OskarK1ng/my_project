package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"registration/internal/models"
	"registration/internal/services"
	"time"
)

type RegisterRequest struct {
	FirstName string    `json:"name"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "method not allowed",
		})

		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("[RegisterHandler] Registering new user: %s %s", user.FirstName, user.LastName)

	if err := services.RegisterUser(context.Background(), &user); err != nil {
		log.Printf("registration failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Registration successful",
		"user_id": user.UserID,
	})
}
