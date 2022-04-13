package services

import (
	"net/http"
	"windows/utility"
)

type ReqStruct struct {
	ServiceName   string `json:"serviceName"`
	ServiceAction string `json:"serviceAction"`
}
type respStruct struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	req := ReqStruct{}
	res := respStruct{
		Status: "failed",
	}
	if err := utility.HttpBodyUnmarshal(r.Body, &req); err != nil {
		res.Msg = err.Error()
		utility.HttpConnectionClose(w, r, http.StatusBadRequest, &res)
		return
	}

	if err := req.execute(); err != nil {
		res.Msg = err.Error()
		utility.HttpConnectionClose(w, r, http.StatusBadRequest, &res)
		return
	}
	res.Status = "ok"
	utility.HttpSendOK(w, r, &res)
}
