package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"transactions/internal/models"
	"transactions/internal/services"
)

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[DepositHandler] Request received")

	// Проверяем метод
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		log.Printf("Invalid method: %s", r.Method)
		return
	}

	// Логируем тело запроса
	body, _ := io.ReadAll(r.Body)
	log.Printf("Request body: %s", string(body))
	r.Body.Close()

	// Декодируем JSON
	var req models.TransactionRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		log.Printf("JSON decode error: %v", err)
		return
	}

	// Проверяем значения
	if req.Amount <= 0 {
		http.Error(w, "amount must be greater than 0", http.StatusBadRequest)
		log.Printf("Invalid amount: %v", req.Amount)
		return
	}

	// Вызываем бизнес-логику
	err := services.Deposit(context.Background(), req.UserID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Service error: %v", err)
		return
	}

	// Возвращаем ответ
	resp := map[string]string{
		"message": fmt.Sprintf("Deposit %.2f to account %v successful", req.Amount, req.UserID),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	log.Printf("Deposit success for ID=%v, amount=%.2f", req.UserID, req.Amount)
}

func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[WithdrawHandler] Request received")

	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		log.Printf("Invalid method: %s", r.Method)
		return
	}

	body, _ := io.ReadAll(r.Body)
	log.Printf("Request body: %s", string(body))
	r.Body.Close()

	var req models.TransactionRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		log.Printf("JSON decode error: %v", err)
		return
	}

	if req.Amount <= 0 {
		http.Error(w, "amount must be greater than 0", http.StatusBadRequest)
		log.Printf("Invalid amount: %v", req.Amount)
		return
	}

	err := services.Withdraw(context.Background(), req.UserID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Service error: %v", err)
		return
	}

	resp := map[string]string{
		"message": fmt.Sprintf("Withdraw %.2f from account %v successful", req.Amount, req.UserID),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	log.Printf("Withdraw success for ID=%v, amount=%.2f", req.UserID, req.Amount)
}
