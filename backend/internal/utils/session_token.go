package utils

import (
	"crypto/rand"
	"encoding/base64"
)

const TokenLength = 32

func GenerateSecureToken() (string, error) {
	// byte slice to hold the random bytes
	randomBytes := make([]byte, TokenLength)

	// filling the slice with random bytes
	_, err := rand.Read(randomBytes)

	if err != nil {
		return "", err
	}

	// encoding the random bytes to url safe base64 string
	// using RawURLEncoding to avoid chars like + and /, which need escaping

	token := base64.RawURLEncoding.EncodeToString(randomBytes)

	return token, nil
}
