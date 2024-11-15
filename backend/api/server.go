package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ginServer *gin.Engine
}

func NewServer() Server {
	s := gin.Default()
	return Server{
		ginServer: s,
	}
}

func (s Server) Start() error {
	if s.ginServer == nil {
		return fmt.Errorf("server does not exist")
	}
	return s.ginServer.Run()
}
