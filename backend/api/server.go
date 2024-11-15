package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
)

type Server struct {
	ginServer *gin.Engine
}

func NewServer() Server {
	gs := gin.Default()

	s := Server{
		ginServer: gs,
	}
	s.register()
	return s
}

// to register a new http call add the path and function here
func (s *Server) register() error {

	s.ginServer.GET("ping", s.ping)
	s.ginServer.POST("create_user", s.createUser)

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
