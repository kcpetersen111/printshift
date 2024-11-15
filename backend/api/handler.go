package api

import (
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
	if dbPass != pass {
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

// still needs the checks to see if this user can create the level of users
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
	password := gjson.GetBytes(req, "password").String()

	_, err = s.db.Exec("Insert into users values ($1, $2, $3, $4);", email, name, level, password)
	if err != nil {
		slog.Error("error inserting into db: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "error inserting new user to db"))
		return
	}
	slog.Info("successfully created user: %v", email)
	c.JSON(http.StatusCreated, "")

}
