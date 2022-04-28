package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProcessHandler(t *testing.T) {
	reqBody := &ReqStruct{
		ServiceName:   "test-service",
		ServiceAction: "restart",
	}

	byteBody, _ := json.Marshal(reqBody)
	breader := bytes.NewBuffer(byteBody)

	req, err := http.NewRequest(http.MethodPost, "/services", breader)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)
	handler.ServeHTTP(rr, req)

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal("error reading response body")
	}
	var bodyModel respStruct
	json.Unmarshal(body, &bodyModel)

	if bodyModel.Status != "ok" {
		if !strings.Contains(bodyModel.Msg, "executable file not found") { //run net command on linux will return this message and it's ok
			t.Fatalf("Response status is not ok msg: %s", bodyModel.Msg)
		}
	}
}
