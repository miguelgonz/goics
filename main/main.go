package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goics/routes"
	_ "log"
	"net/http"
)

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}
func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}
func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}
func (m *mockResponseWriter) WriteHeader(int) {}

const VERSION = "0.1"

func main() {
	router := httprouter.New()
	router.GET("/list/status", routes.Status)
	router.GET("/list/channel/:channel/a-z", routes.ChannelAtoz)

	req, _ := http.NewRequest("GET", "/list/channel/bbc_one_london/a-z", nil)
	router.ServeHTTP(new(mockResponseWriter), req)
	//log.Fatal(http.ListenAndServe(":7080", router))
}
