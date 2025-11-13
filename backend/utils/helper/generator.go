package helper

import (
	"crypto/rand"
	"fmt"
)

func GenerateNim(prefix string, start, end int) []string {
	var users []string
	for i := start; i <= end; i++ {
		users = append(users, fmt.Sprintf("%s%03d", prefix, i))
	}
	return users
}

func GenerateRandomPassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	for i := range b {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b), nil
}
