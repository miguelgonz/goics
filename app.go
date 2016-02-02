package goics

import (
	"github.com/julienschmidt/httprouter"
	"goics/routes"
	"net/http"
)

//go:generate go-bindata -o templates/templates.go -ignore \"\.go$\" -prefix templates/ -pkg templates templates/...

func NewRouter() http.Handler {
	router := httprouter.New()
	router.GET("/status", routes.Status)
	router.GET("/list/channel/:channel/a-z", routes.ChannelAtoz)
	return router
}
