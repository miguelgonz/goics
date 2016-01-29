package routes

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type StatusResponse struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

func Status(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	output, _ := json.Marshal(StatusResponse{
		Version: "0.1",
		Name:    "list_service",
	})
	w.Write(output)
}
