package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) removeUserFromClass(c *gin.Context) {
	var requestBody RemoveUserFromClass

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	_, err := s.db.Exec("delete from class_users where user_id = $1 and class_id = $2;", requestBody.UserId, requestBody.ClassId)
	if err != nil {
		slog.Error("error removing user from class: %v", err)

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func (s *Server) addUserToClass(c *gin.Context) {
	var requestBody AddUserToClass

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	_, err := s.db.Exec("insert into class_users (user_id, class_id) values ($1, $2);", requestBody.UserId, requestBody.ClassId)
	if err != nil {
		slog.Error("error adding user to class: %v", err)

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, "OK")
}

func (s *Server) listPrintersForClass(c *gin.Context) {
	var requestBody ListPrintersRequest

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	rows, err := s.db.Query(`
		select 
			p.id,
			p.name,
			p.description,
			p.is_active
		from printers p
		join class_printers cp
			on cp.printer_id = p.id
		join classes c
			on c.id = cp.class_id
		where c.id = $1
	`, requestBody.ClassId)

	if err != nil {
		slog.Error("error listing classes for user: %v", err)

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

func (s *Server) listClassesForUser(c *gin.Context) {
	var requestBody ListClassesRequest

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	rows, err := s.db.Query(`
		select 
			c.id,
			c.name,
			c.description,
			c.is_active
		from users u
		join class_users cu
			on cu.user_id = u.id
		join classes c
			on c.id = cu.class_id
		where u.id = $1
	`, requestBody.UserId)

	if err != nil {
		slog.Error("error listing classes for user: %v", err)

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	classes := make([]Class, 0)

	for rows.Next() {
		var cl Class
		if err := rows.Scan(&cl.Id, &cl.Name, &cl.Description, &cl.IsActive); err != nil {
			slog.Error("error scanning db: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		classes = append(classes, cl)
	}

	c.JSON(http.StatusOK, classes)
}

func (s *Server) addPrinterToClass(c *gin.Context) {
	var requestBody AddPrinterToClass

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	_, err := s.db.Exec("insert into class_printers (class_id, printer_id) values ($1, $2);", requestBody.ClassId, requestBody.PrinterId)
	if err != nil {
		slog.Error("error adding printer to class: %v", err)

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, "OK")
}
