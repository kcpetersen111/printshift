package main

import (
	"backend/api"
	"backend/persist"
	"log"
)

func main() {
	db := persist.NewDB()
	defer db.Close()

	s := api.NewServer(db)
	if err := s.Start(); err != nil {
		log.Fatalf("error running server: %v", err)
	}
	log.Printf("Server Closing")
}
