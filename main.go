package main

import (
	"windows/aaa"
	"windows/config"

	"github.com/gorilla/mux"
)

func main() {

	// config.LoadMainConfig(config.FileDefaultPath)
	config.FlagParser()
	router := mux.NewRouter()
	aaa.SetEnforcer()
	AddRoutes(router)

	// v1
	AddV1Routes(router.PathPrefix("/v1").Subrouter())
	go listenHTTP(router)
	listenHTTPS(router)

}
