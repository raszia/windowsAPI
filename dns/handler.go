package dns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ReqStruct struct {
	Type       string           `json:"type"` //editRecord,addRecord,removeRecord
	EditRecord EditRecordStruct `json:"editRecord"`
}
type respStruct struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

type EditRecordStruct struct {
	ZoneName   string `json:"zoneName"`   //"myco.local"
	RecordName string `json:"recordName"` //"test"
	RecordType string `json:"recordType"` //"cname"
	RecordData string `json:"recordData"` //"test2.myco.local"
}

//{"type":"editRecord","editRecord":{"zoneName":"my-co.ir","recordName":"test","recordType":"cname","recordData":"test.test.test4"}}
func Handler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req := &ReqStruct{}
	if err := json.Unmarshal(body, req); err != nil {
		fmt.Println(err)
		return
	}
	resp := requestDispatch(req)
	respBody, _ := json.Marshal(resp)
	if resp.Status == "failed" {
		w.Header().Add("connection", "close")
		w.WriteHeader(http.StatusBadGateway)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBody)
}

func requestDispatch(req *ReqStruct) *respStruct {

	resp := &respStruct{
		Status: "failed",
	}
	switch req.Type {
	case "editRecord":
	case "addRecord":
	case "deleteRecord":
	default:
		resp.Msg = "bad req type"
		return resp
	}
	err := req.EditRecord.execute(req.Type)
	if err != nil {
		resp.Msg = err.Error()
		return resp
	}

	resp.Status = "ok"
	return resp
}
