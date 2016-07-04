package server

import (
	"errors"
	"fmt"

	"github.com/afrostream/afrostream-go-lib/env"
	"github.com/afrostream/afrostream-go-lib/log"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   uint16
	host   string
	Engine *gin.Engine
}

func (server *Server) Spawn(port uint16) error {
	var addr string
	var envPort uint16
	var err error

	err, envPort = env.ReadAsUint16("PORT")
	if port != 0 {
		server.port = port
	} else if envPort != 0 && err == nil {
		server.port = envPort
	} else {
		log.Log(log.ERROR, "missing port")
		return errors.New("missing port")
	}
	// addr = ":port"
	addr = fmt.Sprintf(":%d", server.port)
	log.Log(log.INFO, "Starting afrostream-api on %s", addr)
	server.Engine.Run(addr)
	return nil
}

func New() *Server {
	var server *Server

	server = new(Server)
	server.Engine = gin.New()
	return server
}
