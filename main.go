package main

import (
	"sync"
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

	wg := new(sync.WaitGroup)
	if config.Main().WebServer().DisableHttp == "false" {
		wg.Add(1)
		go listenHTTP(router, wg)
	}
	if config.Main().WebServer().DisableHttps == "false" {
		wg.Add(1)
		go listenHTTPS(router, wg)
	}
	wg.Wait()
}
