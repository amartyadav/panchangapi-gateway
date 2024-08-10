package models

import "time"

type User struct {
	ID int `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email string `json:"email": db:"email"`
	PasswordHash string `json:"-" db:"password_hash"`
	APIKey string `json:"api_key" db:"api_key"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	LastLogin time.Time `json:"last_login" db:"last_login"`
	Status string `json:"status" db:"status"`
}
