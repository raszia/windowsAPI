package dns

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"windows/utility"

	"github.com/gorilla/mux"
)

func infoHandler(w http.ResponseWriter, r *http.Request) {
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

func zonesHandler(w http.ResponseWriter, r *http.Request) {
	info, err := getZones(r.Context())
	if err != nil {
		utility.HttpConnectionClose(w, r, http.StatusNotAcceptable, &utility.ResStruct{
			Msg:    err.Error(),
			Status: "failed",
		})
		return
	}
	utility.HttpSendOKByte(w, r, info)
}

type respStruct struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

//{"recordName":"test","recordType":"cname","recordData":"test.test.test4"}
func recordActionHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	req := &recordStruct{
		ZoneName:   vars["zoneName"],
		RecordName: vars["recordName"],
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.HttpConnectionClose(w, r, http.StatusNotAcceptable, &utility.ResStruct{
			Msg:    err.Error(),
			Status: "failed",
		})
		return
	}

	json.Unmarshal(body, req)

	err = recordAction(r.Context(), req, r.Method)
	if err != nil {
		utility.HttpConnectionClose(w, r, http.StatusNotAcceptable, &utility.ResStruct{
			Msg:    err.Error(),
			Status: "failed",
		})
		return
	}
	utility.HttpSendOK(w, r, &respStruct{
		Status: "ok",
		Msg:    "successful",
	})
}
