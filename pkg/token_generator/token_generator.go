package token_generator

import (
	"math/rand"
	"time"
)

const (
	Numbers     = "0123456789"
	Alphabets   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Base64Token = Numbers + Alphabets + "-._~+/"
)

// GenerateRandomNumStr generates digit random number string
func GenerateRandomNumStr(digit uint) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, digit)
	for i := range bytes {
		bytes[i] = Numbers[rand.Intn(len(Numbers))]
	}
	return string(bytes)
}

// GenerateRandomStr generates digit random string
func GenerateRandomBase64Token(digit uint) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, digit)
	for i := range bytes {
		bytes[i] = Base64Token[rand.Intn(len(Base64Token))]
	}
	return string(bytes)
}
