package main

import (
	"log"
	"net/http"
	"time"
	"windows/config"

	"github.com/gorilla/mux"
)

func listenHTTP(router *mux.Router) {
	httpServer := &http.Server{
		Addr:           config.MainConfig.WebServerS.HttpAddr,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("HTTP server listen:", httpServer.Addr)
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
		Addr:           config.MainConfig.WebServerS.HttpsAddr,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("HTTPS server listen:", httpServer.Addr)
	log.Fatal(httpServer.ListenAndServeTLS(config.MainConfig.WebServerS.Cert, config.MainConfig.WebServerS.Key))
}
