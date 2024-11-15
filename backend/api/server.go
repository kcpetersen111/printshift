package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
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

	return nil
}

func (s *Server) Start() error {
	if s.ginServer == nil {
		return fmt.Errorf("server does not exist")
	}
	return s.ginServer.Run()
}
