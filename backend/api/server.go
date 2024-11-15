package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/tidwall/sjson"
)

type Server struct {
	ginServer *gin.Engine
	db        *sql.DB
}

var secretKey = []byte("your-secret-key")

func NewServer(db *sql.DB) Server {
	gs := gin.Default()

	s := Server{
		ginServer: gs,
		db:        db,
	}
	s.register()
	return s
}

// to register a new http call add the path and function here
func (s *Server) register() error {

	s.ginServer.GET("ping", s.ping)
	s.ginServer.GET("login", s.login)

	proc := s.ginServer.Group("/protected")

	proc.POST("create_user", s.createUser)

	s.ginServer.Use(authMiddleware())
	return nil
}

func (s *Server) Start() error {
	if s.ginServer == nil {
		return fmt.Errorf("server does not exist")
	}

	return s.ginServer.Run("0.0.0.0:3410")
}

func mustSet(json, key string, value interface{}) string {
	s, err := sjson.Set(json, key, value)
	if err != nil {
		panic("invalid json")
	}
	return s
}

func mustSetBytes(json []byte, key string, value interface{}) []byte {
	s, err := sjson.SetBytes(json, key, value)
	if err != nil {
		panic("invalid byte json")
	}
	return s
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the cookie
		tokenString, err := c.Cookie("auth_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid token"})
			c.Abort()
			return
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Optionally, you can access token claims here
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("username", claims["username"])
		}

		c.Next()
	}
}
