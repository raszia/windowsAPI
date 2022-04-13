package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"windows/dns"
	"windows/services"
	"windows/utility"

	"github.com/gorilla/mux"
)

const (
	serverPort = ":3111"
	CA         = "./cert/mycoServicesCA.pem"
	Key        = "./cert/mycodc-private.pem"
	Cert       = "./cert/mycodc.myco.local.crt"
)

func main() {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(utility.HttpNotFoundHandler)

	caCert, err := ioutil.ReadFile(CA)
	if err != nil {
		log.Fatal(err)
	}
	router.Methods("POST").Path("/dns").HandlerFunc(dns.Handler)
	router.Methods("POST").Path("/services").HandlerFunc(services.Handler)
	router.Methods("GET").Path("/").HandlerFunc(utility.HttpNotFoundHandler)

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{
		RootCAs:   caCertPool,
		ClientCAs: caCertPool,
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = tlsConfig
	httpServer := &http.Server{
		TLSConfig:      tlsConfig,
		Addr:           serverPort,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Listening on port %s.\n", serverPort)

	// Start HTTP Server
	log.Fatal(httpServer.ListenAndServeTLS(Cert, Key))
}
