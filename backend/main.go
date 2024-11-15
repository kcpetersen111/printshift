package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
)

func main() {

	server := gin.Default()
	server.GET("/ping", func(ctx *gin.Context) {
		j, _ := sjson.Set("", "ping", "pong")
		ctx.JSON(http.StatusOK, j)
	})
	server.Run()
}
