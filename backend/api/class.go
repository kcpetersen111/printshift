package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) listClasses(c *gin.Context) {
	rows, err := s.db.Query(`select id, name, description, is_active from classes;`)
	if err != nil {
		slog.Error("error querying db: %v", err)

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

func (s *Server) updateClass(c *gin.Context) {
	var requestBody UpdateClass

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	_, err := s.db.Exec("update classes set name = $1, description = $2, is_active = $3 where id = $4",
		requestBody.Name,
		requestBody.Description,
		requestBody.IsActive,
		requestBody.ClassId,
	)
	if err != nil {
		slog.Error("error updating class: %v", err)

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// attach professor if applicable
	if requestBody.ProfessorId != -1 {
		rows, err := s.db.Query(`select id from classes where name = $1;`, requestBody.Name)
		if err != nil {
			slog.Error("error querying db: %v", err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		var class Class
		for rows.Next() {
			if err := rows.Scan(&class.Id); err != nil {
				slog.Error("error scanning db: %v", err)

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			break
		}

		_, err = s.db.Exec("insert into class_users (user_id, class_id) values ($1, $2);", requestBody.ProfessorId, class.Id)
		if err != nil {
			slog.Error("error inserting class into db: %v", err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.JSON(http.StatusAccepted, "OK")
}

func (s *Server) createClass(c *gin.Context) {
	var requestBody CreateClassesRequest

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	_, err := s.db.Exec("insert into classes (name, description, is_active) values ($1, $2, $3);", requestBody.Name, requestBody.Description, requestBody.IsActive)
	if err != nil {
		slog.Error("error inserting class into db: %v", err)

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// attach professor if applicable
	if requestBody.ProfessorId != -1 {
		rows, err := s.db.Query(`select id from classes where name = $1;`, requestBody.Name)
		if err != nil {
			slog.Error("error querying db: %v", err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		var class Class
		for rows.Next() {
			if err := rows.Scan(&class.Id); err != nil {
				slog.Error("error scanning db: %v", err)

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			break
		}

		_, err = s.db.Exec("insert into class_users (user_id, class_id) values ($1, $2);", requestBody.ProfessorId, class.Id)
		if err != nil {
			slog.Error("error inserting class into db: %v", err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.JSON(http.StatusCreated, "OK")
}
