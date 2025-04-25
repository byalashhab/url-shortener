package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func GenerateHash(longURL string) (string, int) {
	base62Chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	id := generateUniqueID()

	base := len(base62Chars)
	result := ""

	for id > 0 {
		remainder := id % base
		result = string(base62Chars[remainder]) + result
		id /= base
	}

	return result, id
}

// between 0 - 62^(7)
func generateUniqueID() int {

	str := fmt.Sprintf("%v%v", time.Now().Minute(), rand.Int31())

	num, err := strconv.Atoi(str)

	if err != nil {
		log.Fatal(err)
	}

	return num
}
