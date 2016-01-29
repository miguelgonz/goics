package presenters

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"iblclient"
	"strings"
	"testing"
)

func TestEpisodeHasMarkup(t *testing.T) {
	crawler := getCrawler(getTestItem(), t)
	assertNodeCount(2, ".title", crawler, t)
	assertNodeCount(1, ".primary .overlay", crawler, t)
	assertNodeCount(1, ".primary .play-icon", crawler, t)
	assertNodeCount(1, ".secondary .synopsis", crawler, t)
	assertNodeCount(1, ".secondary .master-brand", crawler, t)

	assertNodeText("The programme title", ".title.top-title", crawler, t)
	assertNodeText("The programme subtitle", ".subtitle", crawler, t)
}

func assertNodeCount(expected int, selector string, crawler *goquery.Document, t *testing.T) {
	foundCount := crawler.Find(selector).Length()
	if foundCount != expected {
		t.Error(fmt.Sprintf("The selector %s was found %d times, expecting %d.", selector, foundCount, expected))
	}
}

func assertNodeText(expected string, selector string, crawler *goquery.Document, t *testing.T) {
	foundText := crawler.Find(selector).Text()
	if foundText != expected {
		t.Error("The selector " + selector + " contains " + foundText + ", expecting " + expected)
	}
}

func getTestItem() (item iblclient.Item) {
	item.Image = iblclient.Image{"http://www.imagechef.com/someimage.jpg"}
	item.Title = "The programme title"
	item.Subtitle = "The programme subtitle"
	item.ID = "p0123456"
	return
}

func getCrawler(item iblclient.Item, t *testing.T) *goquery.Document {
	markup := PresentEpisode(item)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(markup))
	if err != nil {
		t.Error("Invalid markup")
	}
	return doc
}
