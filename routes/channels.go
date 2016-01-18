package routes

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goics/presenters"
	"iblclient"
	"net/http"
)

func ChannelAtoz(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var response string
	url := "channels/" + params.ByName("channel") + "/programmes"
	reqParams := iblclient.RequestParams{}

	response = <-iblclient.Request(url, reqParams)

	res := iblclient.ParseResponse(response)

	firstEpisode := res.Data.Programmes[0]

	presentedFirstEpisode := presenters.PresentEpisode(firstEpisode)

	w.Write([]byte(fmt.Sprintf("%s", presentedFirstEpisode)))
}
