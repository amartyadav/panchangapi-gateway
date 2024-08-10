package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"panchangapi-gateway/internal/database"
	"panchangapi-gateway/internal/utils"
)

type UserRegistrationRequest struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required, min=8"`
}

func UserRegistrationHandler(context *gin.Context) {
	var req UserRegistrationRequest

	if err := context.ShouldBindJSON(&req);

	err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM panchangapiusers WHERE username = $1 OR email = $2)",
		req.Username, req.Email).Scan(&exists)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Database Error"})
		return
	}
	if exists {
		context.JSON(http.StatusConflict, gin.H{"error": "Username or Email aready exists"})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H {"error": "Error hashing password"})
		return
	}

	apiKey := utils.GenerateAPIKey()

	_, err = database.DB.Exec(
		`INSERT INTO panchangapiusers 
		(username, email, password_hash, api_key, status)
		VALUES ($1, $2, $3, $4, $5)`,
	req.Username, req.Email, hashedPassword, apiKey, "active")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H {"error": "Unable to insert user details in db"})
		return
	}

	context.JSON(http.StatusCreated, gin.H {"success": "User created", "api_key": apiKey })
}