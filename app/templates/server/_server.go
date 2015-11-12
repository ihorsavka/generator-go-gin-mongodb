package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	CFG    *Config
	Engine *gin.Engine
}

func New(config *Config) *HTTPServer {
	return &HTTPServer{config, gin.Default()}
}

func (server *HTTPServer) Start() {
	server.Engine.Run(fmt.Sprintf("0.0.0.0:%d", server.CFG.Port))
}
