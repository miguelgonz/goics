package presenters

import (
	"bytes"
	"fmt"
	"goics/templates"
	"html/template"
	"iblclient"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

type templateData struct {
	Number int
	Image  string
	iblclient.Item
}

func PresentEpisode(item iblclient.Item) string {

	episodeTemplate, err := template.New("episode").Parse(string(templates.MustAsset("episode.html")))
	checkError(err)

	var doc bytes.Buffer
	err = episodeTemplate.Execute(&doc, templateData{
		Number: 0,
		Image:  strings.Replace(item.Image.Standard, "{recipe}", "336x189", 1),
		Item:   item,
	})
	checkError(err)

	return doc.String()
}
