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
	if req.Balance <= 0 {
		http.Error(w, "amount must be greater than 0", http.StatusBadRequest)
		log.Printf("Invalid amount: %v", req.Balance)
		return
	}

	// Вызываем бизнес-логику
	err := services.Deposit(context.Background(), req.UserID, req.Balance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Service error: %v", err)
		return
	}

	// Возвращаем ответ
	resp := map[string]string{
		"message": fmt.Sprintf("Deposit %.2f to account %v successful", req.Balance, req.UserID),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	log.Printf("Deposit success for ID=%v, amount=%.2f", req.UserID, req.Balance)
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

	if req.Balance <= 0 {
		http.Error(w, "amount must be greater than 0", http.StatusBadRequest)
		log.Printf("Invalid amount: %v", req.Balance)
		return
	}

	err := services.Withdraw(context.Background(), req.UserID, req.Balance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Service error: %v", err)
		return
	}

	resp := map[string]string{
		"message": fmt.Sprintf("Withdraw %.2f from account %v successful", req.Balance, req.UserID),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	log.Printf("Withdraw success for ID=%v, amount=%.2f", req.UserID, req.Balance)
}

func GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[GetBalanceHandler] Requset received")

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusBadRequest)
		log.Printf("Invalid method: %s", r.Method)
		return
	}

	// Читаем тело запроса
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()
	log.Printf("Request body: %s", string(body))

	var req struct {
		UserID string `json:"user_id"`
	}

	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		log.Printf("JSON decode error: %v", err)
		return
	}

	if req.UserID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		log.Println("user_id is empty")
		return
	}

	// Вызов сервиса
	balance, err := services.GetBalance(context.Background(), req.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Service error: %v", err)
		return
	}

	// Ответ
	resp := map[string]interface{}{
		"user_id": req.UserID,
		"balance": balance,
	}

	json.NewEncoder(w).Encode(resp)

	log.Printf("Balance checked for ID=%v: %.2f", req.UserID, balance)
}
