package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Index(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Afrostream API V2"))
}

func main() {
	var server *Server
	var router *mux.Router

	server = NewServer()
	router = mux.NewRouter()
	router.HandleFunc("/", Index)
	server.Spawn(CONF_DEFAULT_PORT, router)
}
