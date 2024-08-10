package utils

import (
	"context"
	"fmt"
	"time"

	"panchangapi-gateway/internal/database"

	"github.com/redis/go-redis/v9"
)

func StoreOtp(email, otp string) error {
	redisClient := database.GetRedisClient()
	fmt.Printf("[INFO] Email: %s, OTP: %s", email, otp)
	return redisClient.Set(context.Background(), "otp:"+email, otp, 15*time.Minute).Err()
}

func VerifyOtp(email string, otp string) (bool, error) {
	redisClient := database.GetRedisClient()
	storedOtp, err := redisClient.Get(context.Background(), "otp:"+email).Result()

	if err == redis.Nil {
		fmt.Println("[INFO] OTP for this email not found") // OTP not found or expired
		return false, nil
	} else if err != nil {
		return false, err
	}

	return storedOtp == otp, nil
}
