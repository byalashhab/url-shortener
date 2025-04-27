package main

import (
	"math/rand"
)

func GenerateHash(longURL string) (string, int) {
	base62Chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	id := generateUniqueID()

	base := len(base62Chars)
	result := ""

	tempID := id

	for tempID > 0 {
		remainder := tempID % base
		result = string(base62Chars[remainder]) + result
		tempID /= base
	}

	return result, id
}

// between 0 - 62^(7)
func generateUniqueID() int {
	const max = 3_000_000_000_000
    return rand.Intn(int(max))  
}
