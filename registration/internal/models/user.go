package models

import "time"

// Модель пользователя (то, что хранится в БД)
type User struct {
	UserID        string    `json:"user_id"`
	UserName      string    `json:"user_name"`
	UserLastName  string    `json:"user_last_name"`
	UserEmail     string    `json:"user_email"`
	UserPhone     string    `json:"user_phone"`
	UserPassword  string    `json:"user_password"`
	UserCreatedAt time.Time `json:"user_created_at"`
}

// Структура для входящего запроса (от клиента)
type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}
