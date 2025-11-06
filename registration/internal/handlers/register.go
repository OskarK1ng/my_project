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

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("[RegisterHandler] Registering new user: %s %s", req.FirstName, req.LastName)

	user := models.User{
		UserName:      req.FirstName,
		UserLastName:  req.LastName,
		UserEmail:     req.Email,
		UserPhone:     req.Phone,
		UserPassword:  req.Password,
		UserCreatedAt: time.Now(),
	}

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
