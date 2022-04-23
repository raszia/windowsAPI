package main

import (
	"windows/config"

	"github.com/gorilla/mux"
)

func main() {

	// config.LoadMainConfig(config.FileDefaultPath)
	config.FlagParser()
	router := mux.NewRouter()
	AddRoutes(router)

	// v1
	AddV1Routes(router.PathPrefix("/v1").Subrouter())
	go listenHTTP(router)
	listenHTTPS(router)

}
