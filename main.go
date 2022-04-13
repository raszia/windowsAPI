package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/myco/dns"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Could not get the requested route."})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

const (
	serverPort = ":3111"
	CA         = "./cert/mycoServicesCA.pem"
	Key        = "./cert/mycodc-private.pem"
	Cert       = "./cert/mycodc.myco.local.crt"
)

func main() {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	caCert, err := ioutil.ReadFile(CA)
	if err != nil {
		log.Fatal(err)
	}
	router.Methods("POST").Path("/dns").HandlerFunc(dns.Handler)
	router.Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		respondWithJSON(w, http.StatusOK, map[string]string{"message": "Welcome to myco windows API service"})
	})

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
