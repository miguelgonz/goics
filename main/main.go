package main

import (
	_ "fmt"
	"goics"
	"log"
	"net/http"
	_ "net/http/httptest"
)

func main() {
	router := goics.NewRouter()
	log.Fatal(http.ListenAndServe(":7080", router))

	/*
		req, _ := http.NewRequest("GET", "/list/channel/bbc_one_london/a-z", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		fmt.Println(w.Code)
		fmt.Println(w.Body.String())
	*/

}
