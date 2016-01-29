package routes

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goics/presenters"
	"iblclient"
	"net/http"
)

func ChannelAtoz(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	channelProgrammes := iblclient.FetchChannelProgrammes(params.ByName("channel"))
	var presentedItem string

	if len(channelProgrammes.Elements) == 0 {
		w.Write([]byte(fmt.Sprintf("No Episodes: %+v", channelProgrammes)))
	}

	for _, item := range channelProgrammes.Elements {
		presentedItem = presenters.PresentEpisode(item)
		w.Write([]byte(fmt.Sprintf("%s", presentedItem)))
	}
}
