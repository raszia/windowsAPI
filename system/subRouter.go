package system

import (
	"net/http"
	"windows/utility"

	"github.com/gorilla/mux"
)

func SubRoute(router *mux.Router) {
	router.HandleFunc("/info", infoHandler).Methods(http.MethodGet)
	// router.HandleFunc("/vmem", memHandler).Methods(http.MethodGet)
	router.HandleFunc("/process", processHandler).Methods(http.MethodGet)
	// router.HandleFunc("/info", getInfoHandler).Methods(http.MethodGet)
	// router.HandleFunc("/info", getInfoHandler).Methods(http.MethodGet)
}

// func memHandler(w http.ResponseWriter, r *http.Request) {

// 	vmem, err := Getvmem(r.Context())
// 	if err != nil {
// 		utility.HttpConnectionClose(w, r, http.StatusNotAcceptable, &utility.ResStruct{
// 			Msg:    err.Error(),
// 			Status: "failed",
// 		})
// 		return
// 	}
// 	utility.HttpSendOK(w, r, vmem)
// }

func infoHandler(w http.ResponseWriter, r *http.Request) {

	info, err := GetInfo(r.Context())
	if err != nil {
		utility.HttpConnectionClose(w, r, http.StatusNotAcceptable, &utility.ResStruct{
			Msg:    err.Error(),
			Status: "failed",
		})
		return
	}
	utility.HttpSendOK(w, r, info)
}
func processHandler(w http.ResponseWriter, r *http.Request) {

	process, err := GetProcess(r.Context())
	if err != nil {
		utility.HttpConnectionClose(w, r, http.StatusNotAcceptable, &utility.ResStruct{
			Msg:    err.Error(),
			Status: "failed",
		})
		return
	}
	utility.HttpSendOK(w, r, process)
}
