package api

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func (s *Server) ping(c *gin.Context) {
	res, _ := sjson.Set("", "ping", "pong")
	c.JSON(http.StatusOK, res)
}

func (s *Server) createUser(c *gin.Context) {
	// add user auth stuff later
	// auth, _ := c.Get("username")

	req, err := io.ReadAll(c.Request.Body)
	if err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "call_failed", "True"))
		return
	}

	name := gjson.GetBytes(req, "name").String()
	level := gjson.GetBytes(req, "access_level").Int()
	email := gjson.GetBytes(req, "email").String()
	id := uuid.New().String()

	_, err = s.db.Exec("Insert into users values (?, ?, ?, ?, ?);", id, email, name, level)
	if err != nil {
		slog.Error("error inserting into db: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "error inserting new user to db"))
		return
	}
	slog.Info("successfully created user: %v", email)
	c.JSON(http.StatusCreated, "")

}
