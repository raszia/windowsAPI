package system

import (
	"net/http"
	"windows/utility"

	"github.com/gorilla/mux"
)

func SubRoute(router *mux.Router) {
	router.HandleFunc("/info", infoHandler).Methods(http.MethodGet)
	router.HandleFunc("/vmem", memHandler).Methods(http.MethodGet)
	router.HandleFunc("/process", processHandler).Methods(http.MethodGet)
	router.HandleFunc("/interface", interfaceHandler).Methods(http.MethodGet)
	router.HandleFunc("/disk/{operation}", diskHandler).Methods(http.MethodGet)
}

func diskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	switch vars["operation"] {
	case "usage":
	case "partitions":
	case "iocounters":
	}
	iStat, err := GetInterfaces(r.Context())
	if err != nil {
		utility.HttpConnectionClose(w, r, http.StatusNotAcceptable, &utility.ResStruct{
			Msg:    err.Error(),
			Status: "failed",
		})
		return
	}
	utility.HttpSendOK(w, r, iStat)
}

func interfaceHandler(w http.ResponseWriter, r *http.Request) {

	iStat, err := GetInterfaces(r.Context())
	if err != nil {
		utility.HttpConnectionClose(w, r, http.StatusNotAcceptable, &utility.ResStruct{
			Msg:    err.Error(),
			Status: "failed",
		})
		return
	}
	utility.HttpSendOK(w, r, iStat)
}

func memHandler(w http.ResponseWriter, r *http.Request) {

	vmem, err := Getvmem(r.Context())
	if err != nil {
		utility.HttpConnectionClose(w, r, http.StatusNotAcceptable, &utility.ResStruct{
			Msg:    err.Error(),
			Status: "failed",
		})
		return
	}
	utility.HttpSendOK(w, r, vmem)
}

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
