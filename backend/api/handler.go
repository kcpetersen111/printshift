package api

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/tidwall/sjson"
)

// do better user auth not unencrypted password comparison
func (s *Server) login(c *gin.Context) {
	var requestBody UserLogin

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	var userResponse UserLoginResponse

	rows, err := s.db.Query("select email, name, access_level, password from users where email = $1;", requestBody.Email)
	if err != nil {
		slog.Error("error querying db: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}
	var dbPass string
	for rows.Next() {
		if err := rows.Scan(&userResponse.Email, &userResponse.Name, &userResponse.AccessLevel, &dbPass); err != nil {
			slog.Error("error scanning db: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "can't insert new printer to db"})
			return
		}
	}

	if requestBody.Password == "" || dbPass != requestBody.Password {
		slog.Error("email %s with password '%s' not found", requestBody.Email, requestBody.Password)
		time.Sleep(time.Second)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "can't login"})
		return
	}

	// Set the JWT token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": requestBody.Email,
		"exp":      time.Now().Add(1 * time.Hour).Unix(), // 1 hour expiration
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	// Set the JWT in a cookie.
	c.SetCookie("auth_token", tokenString, 60*60, "/", "localhost", false, true)
	c.JSON(http.StatusOK, userResponse)
}

func (s *Server) ping(c *gin.Context) {
	res, _ := sjson.Set("", "ping", "pong")
	c.JSON(http.StatusOK, res)
}

func (s *Server) createAvailableClassTime(c *gin.Context) {
	var requestBody CreateAvailableClassTime

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("insert into class_times (start_time, end_time, class_id) values ($1, $2, $3);",
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

func (s *Server) createAvailablePrinterTime(c *gin.Context) {
	var requestBody CreateAvailablePrinterTime

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("insert into printer_times (start_time, end_time, printer_id) values ($1, $2, $3);",
		requestBody.StartTime,
		requestBody.EndTime,
		requestBody.PrinterId,
	)
	if err != nil {
		slog.Error("error creating available time: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "can't create available time"))
		return
	}

	c.JSON(http.StatusCreated, "OK")
}
