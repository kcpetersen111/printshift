package api

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/tidwall/sjson"
)

func (s *Server) ping(c *gin.Context) {
	res, _ := sjson.Set("", "ping", "pong")
	c.JSON(http.StatusOK, res)
}

// do better user auth not unencrypted password comparison
func (s *Server) login(c *gin.Context) {
	// req, err := io.ReadAll(c.Request.Body)
	// if err != nil {
	// 	slog.Error("error reading request body: %v", err)
	// 	c.JSON(http.StatusBadRequest, mustSet("", "call_failed", "True"))
	// 	return
	// }
	// email := gjson.GetBytes(req, "email").String()
	// pass := gjson.GetBytes(req, "password").String()

	var requestBody UserLogin

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	rows, err := s.db.Query("select password from users where email = $1;", requestBody.Email)
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
	if requestBody.Password == "" || dbPass != requestBody.Password {
		slog.Error("email %s with password '%s' not found", requestBody.Email, requestBody.Password)
		time.Sleep(time.Second)
		c.JSON(http.StatusUnauthorized, "")
		return
	}

	// Set the JWT token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": requestBody.Email,
		"exp":      time.Now().Add(1 * time.Hour).Unix(), // 1 hour expiration
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	// Set the JWT in a cookie.
	c.SetCookie("auth_token", tokenString, 60*60, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}
func (s *Server) createAvailableTime(c *gin.Context) {
	var requestBody CreateAvailableTime

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("insert into available_times (start_time, end_time, class_id) values ($1, $2, $3);",
		requestBody.StartTime,
		requestBody.EndTime,
		requestBody.ClassId,
	)
	if err != nil {
		slog.Error("error creating available time: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "can't create available time"))
		return
	}

	c.JSON(http.StatusCreated, "OK")
}
