package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
)

func (s *Server) ping(c *gin.Context) {
	res, _ := sjson.Set("", "ping", "pong")
	c.JSON(http.StatusOK, res)
}
