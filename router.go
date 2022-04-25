package main

import (
	"net/http"
	"windows/aaa"
	"windows/dns"
	"windows/ldap"
	"windows/services"
	"windows/system"
	"windows/utility"

	"github.com/gorilla/mux"
)

// AddRoutes takes a router or subrouter and adds all the latest
// routes to it
func AddRoutes(router *mux.Router) {
	router.NotFoundHandler = http.HandlerFunc(utility.HttpNotFoundHandler)
}

// AddV1Routes takes a router or subrouter and adds all the v1
// routes to it
func AddV1Routes(router *mux.Router) {
	router.Use(aaa.BasicAuthMiddleware)
	router.Methods("POST").Path("/dns").HandlerFunc(dns.Handler)
	router.Methods("POST").Path("/services").HandlerFunc(services.Handler)
	router.Methods("GET").Path("/").HandlerFunc(utility.HttpNotFoundHandler)
	// router.Methods("POST").Path("/ldap/userInfo").HandlerFunc(ldap.Handler)
	ldap.SubRoute(router.PathPrefix("/ldap").Subrouter())
	system.SubRoute(router.PathPrefix("/system").Subrouter())
}
