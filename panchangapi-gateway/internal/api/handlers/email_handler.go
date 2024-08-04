package handlers

import (
	"fmt"
	"net/http"
	"net/smtp"
	"panchangapi-gateway/internal/models"

	email "github.com/jordan-wright/email"
	"github.com/labstack/echo/v4"
	"github.com/nanorand/nanorand"
)

func VerifyEmail(c echo.Context) error {
	var req models.UserEmailVerificationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// var exists bool
	// err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM panchangapiusers WHERE email = $1)",
	// 	req.Email).Scan(&exists)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database Error"})
	// }

	// if exists {
	// 	return c.JSON(http.StatusConflict, map[string]string{"error": "User with this email already exists"})
	// }

	verification_code, err := nanorand.Gen(9)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate verification code"})
	}

	message := fmt.Sprintf("<h1>Fancy heading html</h1><br/><h3>Your verification code is <h1>%s</h1>.</h3>", verification_code)

	e := email.NewEmail()
	e.From = "Amartya Yo <amartyadav@gmail.com>"
	e.To = []string{"amartyadav@live.co.uk"}
	e.Subject = "Test email go"
	e.HTML = []byte(message)

	err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "amartyadav@gmail.com", "vicv qihx nreq uqom", "smtp.gmail.com"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send email"})
	}

	return c.NoContent(http.StatusOK)
}
