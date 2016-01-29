package goics

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusReturns200(t *testing.T) {
	w := getRecorderForURL("/status")
	if w.Code != 200 {
		t.Error("Status route is not ok")
	}
}

func TestChannelReturns200(t *testing.T) {
	w := getRecorderForURL("/list/channel/bbc_one/a-z")
	if w.Code != 200 {
		t.Error("Channel route is not ok")
	}
}

func TestNotChannelReturns404(t *testing.T) {
	w := getRecorderForURL("/list/chanel/bbc_one/a-z")
	if w.Code != 404 {
		t.Error("Status route is not ok")
	}
}

func getRecorderForURL(url string) *httptest.ResponseRecorder {
	router := NewRouter()
	req, _ := http.NewRequest("GET", url, nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	return recorder
}
