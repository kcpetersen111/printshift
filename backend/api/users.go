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
		c.JSON(http.StatusBadRequest, mustSet("", "error", "error inserting new printer to db"))
		return
	}

	_, err := s.db.Exec("Insert into users (email, name, access_level, password) values ($1, $2, $3, $4);", req.Email, req.Name, req.Level, req.Password)
	if err != nil {
		slog.Error("error inserting into db: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "error inserting new user to db"))
		return
	}
	slog.Info("successfully created user: %v", req.Email)
	c.JSON(http.StatusCreated, "OK")

}

func (s *Server) listUsers(c *gin.Context) {
	rows, err := s.db.Query(`select id, email, name, access_level from users;`)
	if err != nil {
		slog.Error("error querying db: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "error inserting new printer to db"))
		return
	}
	users := make([]persist.User, 0)
	for rows.Next() {
		var u persist.User
		if err := rows.Scan(&u.Id, &u.Email, &u.Name, &u.AccessLevel); err != nil {
			slog.Error("error scanning db: %v", err)
			c.JSON(http.StatusBadRequest, mustSet("", "error", "error inserting new printer to db"))
			return
		}
		users = append(users, u)
	}

	c.JSON(http.StatusOK, users)
}
