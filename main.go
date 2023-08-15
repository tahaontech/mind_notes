package main

import (
	"log"

	"github.com/tahaontech/mind_notes/db"
	"github.com/tahaontech/mind_notes/server"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	server.NewServer(database, ":3000")
}
