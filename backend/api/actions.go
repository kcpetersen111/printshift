package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) addUserToClass(c *gin.Context) {
	var requestBody AddUserToClass

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("insert into class_users (user_id, printer_id) values ($1, $2);", requestBody.UserId, requestBody.ClassId)
	if err != nil {
		slog.Error("error adding user to class: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't add user to class"})
		return
	}

	c.JSON(http.StatusCreated, "OK")
}

func (s *Server) removeUserFromClass(c *gin.Context) {
	var requestBody RemoveUserFromClass

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("delete from class_users where class_id = $1 and user_id = $2",
		requestBody.UserId,
		requestBody.ClassId,
	)

	if err != nil {
		slog.Error("error adding user to class: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't add user to class"})
		return
	}

	c.JSON(http.StatusCreated, "OK")
}

func (s *Server) addPrinterToClass(c *gin.Context) {
	var requestBody AddPrinterToClass

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("insert into class_printers (class_id, printer_id) values ($1, $2);", requestBody.ClassId, requestBody.PrinterId)
	if err != nil {
		slog.Error("error adding printer to class: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't add printer to class"})
		return
	}

	c.JSON(http.StatusCreated, "OK")
}
