package utils

import (
	"context"
	"fmt"
	"time"

	"panchangapi-gateway/internal/database"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

func StoreSessionData(sessionToken, email, hashedOTP, status string) error {
	redisClient := database.GetRedisClient()
	err := redisClient.HSet(context.Background(), "session:"+sessionToken, "email", email, "otp", hashedOTP, "status", status).Err()

	if err != nil {
		return err
	}

	return redisClient.Expire(context.Background(), "session:"+sessionToken, 15*time.Minute).Err()
}

func GetSessionData(sessionToken string) (string, string, error) {
	redisClient := database.GetRedisClient()

	data, err := redisClient.HGetAll(context.Background(), "session:"+sessionToken).Result()

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

func VerifyOtp(sessionToken, otp string) (bool, error) {
	redisClient := database.GetRedisClient()

	storedHashedOtp, err := redisClient.HGet(context.Background(), "session:"+sessionToken, "otp").Result()

	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHashedOtp), []byte(otp))

	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
