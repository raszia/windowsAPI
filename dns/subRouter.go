package dns

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SubRoute(router *mux.Router) {

	router.HandleFunc("/info", infoHandler).Methods(http.MethodGet)                               //get DNS server info
	router.HandleFunc("/zones", zonesHandler).Methods(http.MethodGet)                             //get a list of zones
	router.HandleFunc("/{zoneName}/{recordName}", recordActionHandler).Methods(http.MethodDelete) //delete a record from a zone
	router.HandleFunc("/{zoneName}/{recordName}", recordActionHandler).Methods(http.MethodPost)   //create a record in a zone
}
