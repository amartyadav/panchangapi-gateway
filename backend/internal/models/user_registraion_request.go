package models

type UserEmailVerificationRequest struct {
	Email string `json:"email" binding:"required,email"`
}