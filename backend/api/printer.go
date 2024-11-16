package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) createPrinter(c *gin.Context) {
	var requestBody CreatePrinterRequest

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("insert into printers (name, is_active) values ($1, $2);", requestBody.Name, requestBody.IsActive)
	if err != nil {
		slog.Error("error inserting printers into db: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "error inserting new printer to db"))
		return
	}

	c.JSON(http.StatusCreated, "OK")
}

func (s *Server) bookPrinter(c *gin.Context) {
	var requestBody BookPrinter

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("insert into user_bookings (user_id, start_time, end_time, printer_id) values ($1, $2, $3, $4);",
		requestBody.UserId,
		requestBody.StartTime,
		requestBody.EndTime,
		requestBody.PrinterId,
	)
	if err != nil {
		slog.Error("error creating booking: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "can't create booking"))
		return
	}

	c.JSON(http.StatusCreated, "OK")
}
