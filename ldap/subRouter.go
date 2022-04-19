package ldap

import (
	"net/http"
	"windows/utility"

	"github.com/gorilla/mux"
)

type ReqStruct struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	BaseDN    string `json:"baseDN"`
	LimitSize int    `json:"limitSize"`
}

func SubRoute(router *mux.Router) {
	router.HandleFunc("/getUserInfo", getUserInfoHandler).Methods(http.MethodPost)
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	req := &ReqStruct{}
	if err := utility.HttpBodyUnmarshal(r.Body, req); err != nil {
		utility.HttpConnectionClose(w, r, http.StatusBadRequest, &utility.ResStruct{
			Msg:    err.Error(),
			Status: "failed",
		})
		return
	}

	ldapConn, err := req.bindConnection()
	if err != nil {
		utility.HttpConnectionClose(w, r, http.StatusNotAcceptable, &utility.ResStruct{
			Msg:    err.Error(),
			Status: "failed",
		})
		return
	}

	res, err := ldapConn.getUserInfoLdap()
	if err != nil {
		utility.HttpConnectionClose(w, r, http.StatusNotAcceptable, &utility.ResStruct{
			Msg:    err.Error(),
			Status: "failed",
		})
		return
	}
	res.Status = "ok"
	utility.HttpSendOK(w, r, &res)
}
