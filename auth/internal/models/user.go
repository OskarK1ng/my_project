package models

type User struct {
	UserID           string `json:"user_id"`
	UserEmail        string `json:"user_email"`
	UserPasswordHash string `json:"user_password_hash"`
}
