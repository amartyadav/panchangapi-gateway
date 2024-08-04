package utils

import (
	"github.com/google/uuid"
)

func GenerateAPIKey() string {
	return "pk_" + uuid.NewString()
}
