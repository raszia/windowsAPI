package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	HTTPSAddrDefault = ":8443"
	HTTPAddrDefault  = ":8080"
	CA               = "./cert/CA.pem"
	Key              = "./cert/ss.key"
	Cert             = "./cert/ss.crt"
)

func listenHTTP(router *mux.Router) {
	httpServer := &http.Server{
		Addr:           HTTPAddr,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("HTTP listening on Addr:", HTTPAddr)
	log.Fatal(httpServer.ListenAndServe())

}

func listenHTTPS(router *mux.Router) {
	// caCert, err := ioutil.ReadFile(CA)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(caCert)
	// tlsConfig := &tls.Config{
	// 	RootCAs:   caCertPool,
	// 	ClientCAs: caCertPool,
	// }
	// http.DefaultTransport.(*http.Transport).TLSClientConfig = tlsConfig
	httpServer := &http.Server{
		// TLSConfig:      tlsConfig,
		Addr:           HTTPSAddr,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("HTTPS listening on Addr:", HTTPSAddr)

	//Start HTTP Server
	log.Fatal(httpServer.ListenAndServeTLS(Cert, Key))
}
