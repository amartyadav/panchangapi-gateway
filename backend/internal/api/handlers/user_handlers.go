package handlers

import (
	"net/http"
	"panchangapi-gateway/internal/database"
	"panchangapi-gateway/internal/utils"

	"github.com/labstack/echo/v4"
)

type UserRegistrationRequest struct {
	SessionToken string `json:"sessionToken" binding:"required"`
	Password     string `json:"password" binding:"required, min=8"`
}

func CreateProfile(c echo.Context) error {
	var req UserRegistrationRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	email, status, err := utils.GetSessionData(req.SessionToken)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid session"})
	}

	if status != "verified" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email not verified"})
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error hashing password"})
	}

	apiKey := utils.GenerateAPIKey()

	_, err = database.DB.Exec(
		`INSERT INTO users
		(email, password_hash, api_key)
		VALUES ($1, $2, $3)`,
		email, hashedPassword, apiKey)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"success": "User created", "api_key": apiKey})
}
