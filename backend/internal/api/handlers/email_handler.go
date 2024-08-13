package handlers

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"panchangapi-gateway/internal/database"
	"panchangapi-gateway/internal/models"
	"panchangapi-gateway/internal/utils"

	email "github.com/jordan-wright/email"
	"github.com/labstack/echo/v4"
	"github.com/nanorand/nanorand"
)

func VerifyEmail(c echo.Context) error {
	utils.LoadEnv()

	password := os.Getenv("GMAIL_APP_SPECIFIC_PASSWORD")

	var req models.UserEmailVerificationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)",
		req.Email).Scan(&exists)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if exists {
		return c.JSON(http.StatusConflict, map[string]string{"error": "User with this email already exists"})
	}

	isSignupBlockedForEmail, signupAttemptErr := utils.IsSignupAttemptBlocked(req.Email)

	if signupAttemptErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": signupAttemptErr.Error()})
	}

	if isSignupBlockedForEmail == true {
		return c.JSON(http.StatusTooManyRequests, map[string]string{"error": "Too many sign-up requests.\nBlocked for some time.\nTry again later."})
	} else {
		verification_code, err := nanorand.Gen(9)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate verification code"})
		}

		sessionToken, error := utils.GenerateSecureToken()
		if error != nil {
			return err
		}

		err = utils.StoreSessionData(sessionToken, req.Email, utils.HashOTP(verification_code), "initiated")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to store session data"})
		}

		message := fmt.Sprintf("<h1>PanchangAPI Verificaiton</h1><br/><h3>Your verification code is <h1>%s</h1>.</h3>", verification_code)

		e := email.NewEmail()
		e.From = "PanchangAPI <amartyadav@gmail.com>"
		e.To = []string{req.Email}
		e.Subject = "Verification Code - PanchangAPI"
		e.HTML = []byte(message)

		err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "amartyadav@gmail.com", password, "smtp.gmail.com"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send email"})
		}

		utils.AddSignupAttempt(req.Email)

		return c.JSON(http.StatusOK, map[string]string{"session": sessionToken})
	}
}

func VerifyOtp(c echo.Context) error {
	utils.LoadEnv()

	var req models.UserOtpVerificationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	email, status, err := utils.GetSessionData(req.SessionToken)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid session"})
	}

	if status != "initiated" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid session status"})
	}

	verified, err := utils.VerifyOtp(email, req.Otp)

	if err != nil {
		fmt.Println("[ERROR] Error verifying OTP: ", err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error verifying OTP"})
	}

	if verified {
		err := utils.UpdateSessionStatus(req.SessionToken, "verified")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "OTP verified but error in saving session status"})
		}
		return c.JSON(http.StatusOK, map[string]string{"status": "success"})
	} else {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid OTP"})
	}
}
