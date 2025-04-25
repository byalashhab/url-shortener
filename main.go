package main

import (
	"fmt"
	"log"
)

func main() {
	db, err := NewStorage()

	if err != nil {
		log.Fatal(err)
	}

    fmt.Printf("%+v",db)

}
