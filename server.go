package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Server struct {
	port uint16
	host string
}

func getEnvPort() uint16 {
	var envPort string
	var port64 uint64
	var port uint16
	var err error

	envPort = os.Getenv("PORT")
	if envPort == "" {
		return 0
	}
	port64, err = strconv.ParseUint(envPort, 10, 16)
	port = uint16(port64)
	if err != nil {
		Log(ERROR, "env.PORT is not an uint16")
		return 0
	}
	return port
}

func (server *Server) Spawn(port uint16, router *mux.Router) error {
	var addr string
	var envPort uint16

	envPort = getEnvPort()
	if port != 0 {
		server.port = port
	} else if envPort != 0 {
		server.port = envPort
	} else {
		Log(ERROR, "missing port")
		return errors.New("missing port")
	}
	addr = fmt.Sprintf(":%d", server.port)
	Log(INFO, "Starting afrostream-api on %s", addr)
	http.ListenAndServe(addr, router)
	return nil
}

func NewServer() *Server {
	return new(Server)
}
