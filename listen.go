package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"windows/config"

	"github.com/gorilla/mux"
)

type serverStruct struct {
	httpServer  *http.Server
	shutdownReq chan bool
	reqCount    uint32
}

func createNewServer(httpServer *http.Server) *serverStruct {
	//create server
	return &serverStruct{
		httpServer:  httpServer,
		shutdownReq: make(chan bool),
	}
}

func (server *serverStruct) WaitShutdown(stopLog string) {
	irqSig := make(chan os.Signal, 1)
	signal.Notify(irqSig, syscall.SIGINT, syscall.SIGTERM)

	//Wait interrupt or shutdown request through /shutdown
	select {
	case sig := <-irqSig:
		log.Printf("Shutdown request (signal: %v)", sig)
	case sig := <-server.shutdownReq:
		log.Printf("Shutdown request (/shutdown %v)", sig)
	}

	log.Println(stopLog)

	//Create shutdown context with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//shutdown the server
	err := server.httpServer.Shutdown(ctx)
	if err != nil {
		log.Printf("Shutdown request error: %v", err)
	}
}

func listenHTTP(router *mux.Router) {

	httpServer := &http.Server{
		Addr:           config.MainConfig.WebServerS.HttpAddr,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server := createNewServer(httpServer)

	done := make(chan bool)
	go func() {
		log.Println("HTTP server listen:", httpServer.Addr)
		err := server.httpServer.ListenAndServe()
		if err != nil {
			log.Printf("Listen and serve: %v", err)
		}
		done <- true
	}()

	//wait shutdown
	server.WaitShutdown("HTTP Server shutting down Addr: " + httpServer.Addr)

	<-done
	log.Printf("HTTP server shutdown completed Addr: " + httpServer.Addr)

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
	server := createNewServer(httpServer)

	done := make(chan bool)
	go func() {
		log.Println("HTTPS server listen:", httpServer.Addr)
		err := server.httpServer.ListenAndServeTLS(config.MainConfig.WebServerS.Cert, config.MainConfig.WebServerS.Key)
		if err != nil {
			log.Printf("Listen and serve TLS: %v", err)
		}
		done <- true
	}()

	//wait shutdown
	server.WaitShutdown("HTTPS Server shutting down Addr: " + httpServer.Addr)

	<-done
	log.Printf("HTTPS server shutdown completed Addr: " + httpServer.Addr)
}
