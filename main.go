package main

import (
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	AddRoutes(router)

	// v1
	AddV1Routes(router.PathPrefix("/v1").Subrouter())
	go listenHTTP(router)
	listenHTTPS(router)

}
