package models

import (
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
