package utility

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type ResStruct struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func HttpBodyUnmarshal(reqBody io.ReadCloser, model interface{}) error {
	body, err := ioutil.ReadAll(reqBody)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}

func HttpConnectionClose(w http.ResponseWriter, r *http.Request, statusCode int, model interface{}) {
	body, _ := json.Marshal(model)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("connection", "close")
	w.WriteHeader(statusCode)
	w.Write(body)
}

func HttpSendOK(w http.ResponseWriter, r *http.Request, model interface{}) {
	body, _ := json.Marshal(model)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func HttpNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := json.Marshal(map[string]string{"status": "failed", "msg": "Could not get the requested route."})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(body)
}
