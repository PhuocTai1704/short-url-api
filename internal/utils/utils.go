package utils

import (
	"crypto/sha1"
	"os"
)


func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

// Base62 charset
var base62chars = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// DeterministicShort generates the same short code for the same URL
func DeterministicShort(url string, length int) string {
	h := sha1.Sum([]byte(url)) // SHA1 hash URL
	code := ""
	for i := 0; i < length; i++ {
		code += string(base62chars[h[i]%62])
	}
	return code
}


