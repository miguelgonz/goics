package main

import (
	"fmt"
	"goics"
	"iblclient"
	"log"
	"net/http"
	"net/http/httptest"
)

func start() {
	log.Fatal(http.ListenAndServe(":7080", goics.NewRouter()))
}

func testResponse() {
	req, _ := http.NewRequest("GET", "/list/channel/bbc_one_london/a-z", nil)
	w := httptest.NewRecorder()

	goics.NewRouter().ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body.String())
}

func testIblFetchers() {
	fmt.Printf("%+v", iblclient.FetchChannels())
	fmt.Printf("%+v", iblclient.FetchSearch("eastenders"))
}

func main() {
	start()
	//testIblFetchers();
	//testResponse();
}
