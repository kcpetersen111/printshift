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

func (s *Server) addUserToClass(c *gin.Context) {
	var requestBody AddUserToClass

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("insert into class_users (user_id, printer_id) values ($1, $2);", requestBody.UserId, requestBody.ClassId)
	if err != nil {
		slog.Error("error adding user to class: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't add user to class"})
		return
	}

	c.JSON(http.StatusCreated, "OK")
}

func (s *Server) removeUserFromClass(c *gin.Context) {
	var requestBody RemoveUserFromClass

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("delete from class_users where class_id = $1 and user_id = $2",
		requestBody.UserId,
		requestBody.ClassId,
	)

	if err != nil {
		slog.Error("error adding user to class: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't add user to class"})
		return
	}

	c.JSON(http.StatusCreated, "OK")
}

func (s *Server) addPrinterToClass(c *gin.Context) {
	var requestBody AddPrinterToClass

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("insert into class_printers (class_id, printer_id) values ($1, $2);", requestBody.ClassId, requestBody.PrinterId)
	if err != nil {
		slog.Error("error adding printer to class: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't add printer to class"})
		return
	}

	c.JSON(http.StatusCreated, "OK")
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

func (s *Server) updateClass(c *gin.Context) {
	var requestBody UpdateClass

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("update classes set name = $1, description = $2, is_active = $3 where id = $4",
		requestBody.Class.Name,
		requestBody.Class.Description,
		requestBody.Class.Active,
		requestBody.ClassId,
	)
	if err != nil {
		slog.Error("error inserting class into db: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "error inserting new printer to db"))
		return
	}

	c.JSON(http.StatusCreated, "OK")
}

func (s *Server) createClass(c *gin.Context) {
	var requestBody CreateClassesRequest

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("insert into classes (name, description, is_active) values ($1, $2, $3);", requestBody.Name, requestBody.Description, requestBody.Active)
	if err != nil {
		slog.Error("error inserting class into db: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "error inserting new printer to db"))
		return
	}

	c.JSON(http.StatusCreated, "OK")
}

func (s *Server) bookPrinter(c *gin.Context) {
	var requestBody BookPrinter

	if err := c.BindJSON(&requestBody); err != nil {
		slog.Error("error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"call_failed": true})
		return
	}

	_, err := s.db.Exec("insert into user_bookings (user_id, start_time, end_time, printer_id) values ($1, $2, $3, $4);",
		requestBody.UserId,
		requestBody.StartTime,
		requestBody.EndTime,
		requestBody.PrinterId,
	)
	if err != nil {
		slog.Error("error creating booking: %v", err)
		c.JSON(http.StatusBadRequest, mustSet("", "error", "can't create booking"))
		return
	}

	c.JSON(http.StatusCreated, "OK")
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
