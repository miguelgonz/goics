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

	if len(channelProgrammes.Elements) > 0 {
		firstEpisode := channelProgrammes.Elements[0]
		presentedFirstEpisode := presenters.PresentEpisode(firstEpisode)
		w.Write([]byte(fmt.Sprintf("%s", presentedFirstEpisode)))
	} else {
		w.Write([]byte(fmt.Sprintf("No Episodes: %+v", channelProgrammes)))
	}

}
