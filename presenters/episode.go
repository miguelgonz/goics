package presenters

import (
	"bytes"
	"fmt"
	"html/template"
	"iblclient"
	"os"
)

const episodeTpl = `
<a href="/programmes/{{.Id}}" title="{{.Title}}" class="list-item-link stat" >
    {{if .Number}}
        <div class="number" aria-hidden="true">{{.Number}}</div>
    {{end}}
    <div class="title top-title">{{.Title}}</div>
    <div class="primary">
        <img src=""/>
        <div class="overlay"></div>
        <div class="play-icon">SOMEOTHERTPL</div>
    </div>
    <div class="secondary">
        <div class="title">{{.Title}}</div>
    </div>
</a>
`

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func PresentEpisode(episode iblclient.Programme) string {

	t := template.New("Episode template")
	t, err := t.Parse(episodeTpl)
	checkError(err)

	var doc bytes.Buffer
	err = t.Execute(&doc, episode)
	checkError(err)

	return doc.String()
}
