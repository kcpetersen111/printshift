package main

import (
	"backend/api"
	"log"
)

func main() {

	// server := gin.Default()
	// server.GET("/ping", func(ctx *gin.Context) {
	// 	j, _ := sjson.Set("", "ping", "pong")
	// 	ctx.JSON(http.StatusOK, j)
	// })
	// server.Run()
	s := api.NewServer()
	if err := s.Start(); err != nil {
		log.Fatalf("error running server: %v", err)
	}
	log.Printf("Server Closing")
}
