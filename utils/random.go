package utils

import (
	"math/rand"
	"time"
)

const charset = "0123456789"

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateRandomString generates a random string of the specified length.
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return string(b)
}
