package utils

import (
	"math/rand"
	"time"
)

// Seeding the math random
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate a random shortcode of a specified length
func GenerateShortCode(length int) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortCode := make([]byte, length)

	for i := range shortCode {
		shortCode[i] = characters[rand.Intn(len(characters))]
	}

	return string(shortCode)
}
