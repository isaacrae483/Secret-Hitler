package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var upperLetterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandUpperString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = upperLetterRunes[rand.Intn(len(upperLetterRunes))]
	}
	return string(b)
}
