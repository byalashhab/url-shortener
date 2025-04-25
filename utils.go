package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func GenerateHash(longURL string) string {
	// 7 chars for the hashed value
	// [0-9][a-z][A-Z] 9 + 26 + 26 = 62
	// base 62 hashing

    // id := generateUniqueID()

	return ""
}

// between 0 - 62^(7)
func generateUniqueID() int {

	str := fmt.Sprintf("%v%v", time.Now().Nanosecond(), rand.Int31())

	num, err := strconv.Atoi(str)

	if err != nil {
		log.Fatal(err)
	}

	return num
}
