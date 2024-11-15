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
	auth, _ := c.Get("username")

	req, err := io.ReadAll(c.Request.Body)
	if err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "call_failed", "True"))
		return
	}

	name := gjson.GetBytes(req, "name").String()
	level := gjson.GetBytes(req, "access_level").Int()
	email := gjson.GetBytes(req, "email").String()
	classes := make([]string, 0)
	gjson.GetBytes(req, "classes").ForEach(func(key, value gjson.Result) bool {
		classes = append(classes, value.String())
		return true
	})

}
