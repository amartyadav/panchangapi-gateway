package utils

import (
	"context"
	"fmt"
	"time"

	"panchangapi-gateway/internal/database"

	"golang.org/x/crypto/bcrypt"
)

func StoreSessionData(sessionToken, email, hashedOTP, status string) error {
	redisClient := database.GetRedisClient()

	err := redisClient.HSet(context.Background(), "session:"+sessionToken, "email", email, "otp", hashedOTP, "status", status,).Err()

	if err != nil {
		return err
	}

	return redisClient.Expire(context.Background(), "session:" + sessionToken, 15 * time.Minute).Err()
}

func GetSessionData(sessionToken string) (string, string, error) {
	redisClient := database.GetRedisClient()

	data, err := redisClient.HGetAll(context.Background(), "session:" + sessionToken).Result()

	if err != nil {
		return "", "", err
	}

	email, ok := data["email"]
	if !ok {
		return "", "", fmt.Errorf("[ERROR] Invalid session")
	}

	status, ok := data["status"]
	if !ok {
		return "", "", fmt.Errorf("[ERROR] Invalid session")
	}

	return email, status, nil
}

func UpdateSessionStatus(sessionToken, newStatus string) error {
	redisClient := database.GetRedisClient()

	return redisClient.HSet(context.Background(), "session:"+sessionToken, "status", newStatus).Err()
}

func HashOTP(otp string) string {
	hashedOTP, _ := bcrypt.GenerateFromPassword([]byte(otp), bcrypt.DefaultCost)
	return string(hashedOTP)
}
