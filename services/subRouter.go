package services

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SubRoute(router *mux.Router) {
	router.HandleFunc("/action", actionHandler).Methods(http.MethodPost)
	// router.HandleFunc("/vmem", memHandler).Methods(http.MethodGet)
	// router.HandleFunc("/process", processHandler).Methods(http.MethodGet)
	// router.HandleFunc("/interface", interfaceHandler).Methods(http.MethodGet)
	// router.HandleFunc("/disk/{operation}", diskHandler).Methods(http.MethodGet)
}
