package dns

import (
	"net/http"
	"windows/utility"

	"github.com/gorilla/mux"
)

func SubRoute(router *mux.Router) {

	router.HandleFunc("/info", info).Methods(http.MethodGet)
	router.HandleFunc("/zones", info).Methods(http.MethodGet)
	router.HandleFunc("/{zoneName}", getZone).Methods(http.MethodGet)
	// router.HandleFunc("{zoneName}/{recordName}", getRecordData).Methods(http.MethodGet)
	// router.HandleFunc("{zoneName}/{recordName}", getRecordData).Methods(http.MethodDelete)
}

// func getRecordData(w http.ResponseWriter, r *http.Request) {

// }

func info(w http.ResponseWriter, r *http.Request) {
	info, err := getInfo(r.Context())
	if err != nil {
		utility.HttpConnectionClose(w, r, http.StatusNotAcceptable, &utility.ResStruct{
			Msg:    err.Error(),
			Status: "failed",
		})
		return
	}
	utility.HttpSendOKByte(w, r, info)
}

func zones(w http.ResponseWriter, r *http.Request) {

}
