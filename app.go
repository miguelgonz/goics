package goics

import (
	"github.com/julienschmidt/httprouter"
	"goics/routes"
	"net/http"
)

func NewRouter() http.Handler {
	router := httprouter.New()
	router.GET("/status", routes.Status)
	router.GET("/list/channel/:channel/a-z", routes.ChannelAtoz)
	return router
}
