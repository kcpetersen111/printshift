package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/tidwall/sjson"
)

type Server struct {
	ginServer *gin.Engine
	db        *sql.DB
}

var secretKey = []byte("eW91ci1zZWNyZXQta2V5")

func NewServer(db *sql.DB) Server {
	gs := gin.Default()

	gs.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
	s.ginServer.POST("login", s.login)

	proc := s.ginServer.Group("/protected")

	proc.PATCH("update_user", s.updateUser)
	proc.POST("create_user", s.createUser)
	proc.POST("addUserToClass", s.addUserToClass)

	proc.POST("printer", s.createPrinter)
	proc.POST("addPrinterToClass", s.addPrinterToClass)
	proc.POST("removeUserFromClass", s.removeUserFromClass)
	proc.POST("bookPrinter", s.bookPrinter)

	proc.POST("createAvailableTime", s.createAvailableTime)

	proc.PATCH("class", s.updateClass)
	proc.POST("class", s.createClass)

	proc.GET("list_users", s.listUsers)
	proc.GET("list_classes", s.listClasses)

	proc.Use(authMiddleware())
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
