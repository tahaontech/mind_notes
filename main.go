package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tahaontech/mind_notes/db"
	"github.com/tahaontech/mind_notes/server"
)

func main() {
	// -- setup environment
	setUpFolders()               // folders
	database, err := db.InitDB() //db
	if err != nil {
		log.Fatal(err)
	}

	// start server
	s := server.NewServer(database, ":3000")
	s.Start()
}

func setUpFolders() {
	// create data folder and images inside it
	isExist, err := exists("data")
	if !isExist || err != nil {
		err := os.MkdirAll("data", 0777)
		if err != nil {
			log.Fatal(err)
		}
		err = os.MkdirAll("data/images", 0777)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	isExist, err = exists("data/images")
	if !isExist || err != nil {
		err := os.MkdirAll("data/images", 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	// check UI exists
	isExist, err = exists("UI")
	if !isExist || err != nil {
		log.Fatal("UI folder not founds!")
	}

	fmt.Println("your app ready to run.")
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
