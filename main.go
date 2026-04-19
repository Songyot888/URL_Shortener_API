package main

import (
	"log"
	"my-shortener/model/query"
	"my-shortener/repository"
	"my-shortener/system"
)

func main() {
	// This is the entry point of the application.
	// You can initialize your server and routes here.
	db, err := repository.NewConnectDB()
	if err != nil {
		log.Fatal("❌ Failed to connect to the database")
	}
	log.Println("✅ Successfully connected to the database")

	query.SetDefault(db)
	system.StartServer()
}
