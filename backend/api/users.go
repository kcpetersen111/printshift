package api

import (
	"backend/persist"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// still needs the checks to see if this user can create the level of users
func (s *Server) createUser(c *gin.Context) {
	// add user auth stuff later
	// auth, _ := c.Get("username")
	var req CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		slog.Error("error inserting printers into db: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	_, err := s.db.Exec("Insert into users (email, name, access_level, password) values ($1, $2, $3, $4);", req.Email, req.Name, req.AccessLevel, req.Password)
	if err != nil {
		slog.Error("error inserting into db: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	slog.Info("successfully created user: %v", req.Email)
	c.JSON(http.StatusCreated, "OK")

}

func (s *Server) updateUser(c *gin.Context) {
	// add user auth stuff later
	// auth, _ := c.Get("username")
	var req UpdateUserRequest
	if err := c.BindJSON(&req); err != nil {
		slog.Error("error inserting printers into db: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "name, email, and password are required"})
		return
	}

	_, err := s.db.Exec("update users set email = $1, name = $2, password = $3 where email = $4;", req.Email, req.Name, req.Password, req.Email)
	if err != nil {
		slog.Error("error inserting into db: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusAccepted, "OK")
}

func (s *Server) listProfessors(c *gin.Context) {
	rows, err := s.db.Query(`select id, name from users where access_level = 2;`)
	if err != nil {
		slog.Error("error querying db: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	professors := make([]Professor, 0)
	for rows.Next() {
		var p Professor
		if err := rows.Scan(&p.ProfessorId, &p.Name); err != nil {
			slog.Error("error scanning db: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		professors = append(professors, p)
	}

	c.JSON(http.StatusOK, professors)
}

func (s *Server) listUsers(c *gin.Context) {
	rows, err := s.db.Query(`select id, email, name, access_level from users;`)
	if err != nil {
		slog.Error("error querying db: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	users := make([]persist.User, 0)
	for rows.Next() {
		var u persist.User
		if err := rows.Scan(&u.Id, &u.Email, &u.Name, &u.AccessLevel); err != nil {
			slog.Error("error scanning db: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		users = append(users, u)
	}

	c.JSON(http.StatusOK, users)
}
