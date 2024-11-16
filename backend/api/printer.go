package api

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) updatePrinter(c *gin.Context) {
	var requestBody UpdatePrinterRequest

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	_, err := s.db.Exec("update printers set name = $1, description = $2, is_active = $3 where name = $4;",
		requestBody.Name,
		requestBody.Description,
		requestBody.IsActive,
		requestBody.Name,
	)
	if err != nil {
		slog.Error("error updating printer: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, "OK")
}

func (s *Server) createPrinter(c *gin.Context) {
	var requestBody CreatePrinterRequest

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	_, err := s.db.Exec("insert into printers (name, is_active) values ($1, $2);", requestBody.Name, requestBody.IsActive)
	if err != nil {
		slog.Error("error inserting printers into db: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, "OK")
}

func (s *Server) listPrinters(c *gin.Context) {
	rows, err := s.db.Query(`select id, name, description, is_active from printers;`)
	if err != nil {
		slog.Error("error querying db: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	printers := make([]Printer, 0)
	for rows.Next() {
		var p Printer
		if err := rows.Scan(&p.PrinterId, &p.Name, &p.Description, &p.IsActive); err != nil {
			slog.Error("error scanning db: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		printers = append(printers, p)
	}

	c.JSON(http.StatusOK, printers)
}

func (s *Server) bookSpecificPrinter(c *gin.Context) {
	var req BookSpecificPrinter

	if err := c.BindJSON(&req); err != nil {
		slog.Error("error reading request body: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	tx, err := s.db.Begin()
	if err != nil {
		slog.Error("error starting db tx: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
		if err != nil {
			slog.Error("error committing tx for booking a specific time: %v", err)
		}
	}()

	//validate the user can use that printer
	//
	var rows *sql.Rows
	rows, err = tx.Query("select cp.printer_id from class_printers as cp inner join class_users as cu on cp.class_id = cu.class_id where cu.user_id = $1;", req.UserId)
	if err != nil {
		slog.Error("error querying db: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	matches := false
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			slog.Error("error scanning row: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if id == req.PrinterId {
			matches = true
			break
		}
	}
	if !matches {
		slog.Error("error user does not have access to schedule time on this printer")
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	//validate that the time is open
	rows, err = tx.Query("select count(*) from printer_times where printer_id = $1 and start_time <= $2 and end_time >= $3;", req.PrinterId, req.StartTime, req.EndTime)
	if err != nil {
		slog.Error("error query db: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	//should get 1 row if time is open and 0 results if it is false
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			slog.Error("error scanning row: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}
	if count <= 0 {
		slog.Error("error ")
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	rows, err = tx.Query("select count(*) from class_times where printer_id = $1 and start_time <= $2 and end_time >= $3;", req.PrinterId, req.StartTime, req.EndTime)
	if err != nil {
		slog.Error("error query db: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	//should get 1 row if time is open and 0 results if it is false

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			slog.Error("error scanning row: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}
	if count <= 0 {
		slog.Error("error ")
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	rows, err = tx.Query("select id from user_bookings where printer_id = $1 and start_time >= $2 and end_time <= $3 and start_time >= $4 and end_time <= $5;", req.PrinterId, req.StartTime, req.StartTime, req.EndTime, req.EndTime)
	if err != nil {
		slog.Error("error query db: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			slog.Error("error scanning row: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		//if we get a single record then we fail because someone already has this time slot
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "time is already taken"})
		return
	}

	//schedule printers
	_, err = tx.Exec("insert into user_bookings (user_id, printer_id, start_time, end_time) values ($1,$2,$3,$4);",
		req.UserId,
		req.PrinterId,
		req.StartTime,
		req.EndTime,
	)
	if err != nil {
		slog.Error("error creating booking: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, mustSet("", "error", err))
		return
	}

	c.JSON(http.StatusCreated, "OK")
}
