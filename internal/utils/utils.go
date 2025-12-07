package utils

import (
	"crypto/rand"
	"math/big"
	"net/url"
	"os"
	"strings"
	"time"
)

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

const base62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var base62Len = big.NewInt(int64(len(base62)))

func GenerateShortCode(length int) (string, error) {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, base62Len)
		if err != nil {
			return "", err
		}

		result[i] = base62[num.Int64()]
	}

	return string(result), nil
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

func ParseOptionalDate(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}

	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
