package main

import (
	"flag"

	"github.com/gorilla/mux"
)

var HTTPAddr, HTTPSAddr string

func flagParser() {
	flag.StringVar(&HTTPAddr, "HTTPAddr", HTTPAddrDefault, "http address")
	flag.StringVar(&HTTPSAddr, "HTTPSAddr", HTTPSAddrDefault, "https address")
	flag.Parse()
}

func main() {
	flagParser()
	router := mux.NewRouter()
	AddRoutes(router)

	// v1
	AddV1Routes(router.PathPrefix("/v1").Subrouter())
	go listenHTTP(router)
	listenHTTPS(router)

}
