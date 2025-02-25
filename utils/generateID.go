package utils

import "github.com/google/uuid"

// Generate random id
func GenerateRandomID() string {
	id := uuid.New()
	return id.String()
}
