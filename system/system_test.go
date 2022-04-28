package system

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shirou/gopsutil/v3/host"
)

func TestProcessHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/process", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(processHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal("error reading response body")
	}
	var bodyModel []Process
	json.Unmarshal(body, &bodyModel)

	if len(bodyModel) < 1 {
		t.Fatal("response body is too short")
	}

	for _, process := range bodyModel {
		if process.Pid == 0 {
			t.Fatal("process with pid 0 is not valid")
		}
		if process.Name == "" {
			t.Fatal("process with no name is not valid")
		}
		if len(process.Status) < 1 {
			t.Fatal("process with no status is not valid")
		}
	}
}

func TestInfoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(infoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal("error reading response body")
	}
	var bodyModel host.InfoStat
	json.Unmarshal(body, &bodyModel)

	if bodyModel.Hostname == "" {
		t.Error("no hostname")
	}

	if bodyModel.Uptime == 0 {
		t.Error("uptime is 0")
	}

	if bodyModel.Procs == 0 {
		t.Error("proc is 0")
	}
}
