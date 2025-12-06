package utils

import (
	"crypto/sha1"
	"net/url"
	"os"
	"strings"
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

func ValidateURL(input string) bool {
	domain := os.Getenv("DOMAIN_SHORT")

	input = strings.TrimSpace(input)
	if input == "" {
		return false
	}

	// Không cho phép rút gọn URL của chính hệ thống
	if strings.HasPrefix(input, domain) {
		return false
	}

	u, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	if u.Host == "" {
		return false
	}

	if strings.ContainsAny(input, ` <>"{}|\^`+"`") {
		return false
	}

	return true
}
