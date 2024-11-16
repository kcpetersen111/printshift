package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) updateClass(c *gin.Context) {
	var requestBody UpdateClass

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("update classes set name = $1, description = $2, is_active = $3 where id = $4",
		requestBody.Class.Name,
		requestBody.Class.Description,
		requestBody.Class.IsActive,
		requestBody.ClassId,
	)
	if err != nil {
		slog.Error("error inserting class into db: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "error inserting new printer to db"))
		return
	}

	c.JSON(http.StatusCreated, "OK")
}

func (s *Server) createClass(c *gin.Context) {
	var requestBody CreateClassesRequest

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("insert into classes (name, description, is_active) values ($1, $2, $3);", requestBody.Name, requestBody.Description, requestBody.IsActive)
	if err != nil {
		slog.Error("error inserting class into db: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "error inserting new printer to db"))
		return
	}

	c.JSON(http.StatusCreated, "OK")
}
