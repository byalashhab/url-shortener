package main

import "log"

func main() {
	db, err := NewStorage()

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewServer(":1234", db)
	server.Run()

}
