package api

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func (s *Server) ping(c *gin.Context) {
	res, _ := sjson.Set("", "ping", "pong")
	c.JSON(http.StatusOK, res)
}

func (s *Server) createUser(c *gin.Context) {
	req, err := io.ReadAll(c.Request.Body)
	if err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "call_failed", "True"))
		return
	}
	panic(gjson.GetBytes(req, "name").String())
}
