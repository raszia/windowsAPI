package system

import (
	"net/http"
	"strings"
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

	var (
		res interface{}
		err error
	)
	switch vars["operation"] {
	case "usage":
		res, err = GetDiskUsage(r.Context(), r.URL.Query().Get("path"))
	case "partitions":
		all := r.URL.Query().Get("all")
		if all == "true" {
			res, err = GetDiskPartitions(true)
		} else {
			res, err = GetDiskPartitions(false)
		}

	case "iocounters":
		namesList := strings.Split(r.URL.Query().Get("names"), ",")
		res, err = GetDiskIocounters(namesList...)
	}

	if err != nil {
		utility.HttpConnectionClose(w, r, http.StatusNotAcceptable, &utility.ResStruct{
			Msg:    err.Error(),
			Status: "failed",
		})
		return
	}
	utility.HttpSendOK(w, r, res)
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
