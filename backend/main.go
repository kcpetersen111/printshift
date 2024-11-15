package main

import (
	"backend/api"
	"log"
)

func main() {
	s := api.NewServer()
	if err := s.Start(); err != nil {
		log.Fatalf("error running server: %v", err)
	}
	log.Printf("Server Closing")
}
