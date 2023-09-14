package utils

import (
	"math/rand"
	"time"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

func GenHash(size int) string {
	// seed
	rand.Seed(time.Now().UnixNano())
	// generate hash
	hash := make([]rune, size)
	for i := range hash {
		hash[i] = letters[rand.Intn(len(letters))]
	}
	return string(hash)
}

