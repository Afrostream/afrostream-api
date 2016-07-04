package main

import (
	"net/http"

	"github.com/afrostream/afrostream-go-lib/server"
	"github.com/gin-gonic/gin"
)

func Index(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Afrostream API V2"))
}

func main() {
	var s *server.Server

	s = server.New()
	s.Engine.GET("/hello-world", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	s.Spawn(CONF_DEFAULT_PORT)
}
