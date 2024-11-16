package api

import (
	"backend/persist"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func (s *Server) ping(c *gin.Context) {
	res, _ := sjson.Set("", "ping", "pong")
	c.JSON(http.StatusOK, res)
}

// do better user auth not unencrypted password comparison
func (s *Server) login(c *gin.Context) {

	req, err := io.ReadAll(c.Request.Body)
	if err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "call_failed", "True"))
		return
	}
	email := gjson.GetBytes(req, "email").String()
	pass := gjson.GetBytes(req, "password").String()

	rows, err := s.db.Query("select password from users where email = $1;", email)
	if err != nil {
		slog.Error("error querying db: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "call_failed", "True"))
		return
	}
	var dbPass string
	for rows.Next() {
		if err := rows.Scan(&dbPass); err != nil {
			slog.Error("error scanning db res: %v", err)
			c.JSON(http.StatusBadRequest, mustSet("", "call_failed", "True"))
			return
		}
	}
	if pass != "" || dbPass != pass {
		time.Sleep(time.Second)
		c.JSON(http.StatusUnauthorized, "")
		return
	}

	// Set the JWT token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": email,
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // 1-day expiration
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	// Set the JWT in a cookie.
	c.SetCookie("auth_token", tokenString, 3600*24, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Level    int    `json:"level"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

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

type ListUsers struct {
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

type CreatePrinterRequest struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

func (s *Server) createPrinter(c *gin.Context) {
	var requestBody CreatePrinterRequest

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("insert into printers (name, is_active) values ($1, $2);", requestBody.Name, requestBody.Active)
	if err != nil {
		slog.Error("error inserting printers into db: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "error inserting new printer to db"))
		return
	}

	c.JSON(http.StatusCreated, "OK")
}
