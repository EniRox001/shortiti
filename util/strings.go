package util

import (
	"math/rand"
	"strings"
	"time"
)

func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateUuid(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateShortLink() (string, string) {
	uniqueId := generateUuid(5)
	return "http://localhost:3000/" + uniqueId, uniqueId
}
