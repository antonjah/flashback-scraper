package thread

import (
	"github.com/PuerkitoBio/goquery"
)

func getTitle(doc *goquery.Document) string {
	title := doc.Find("title").Contents().Text()
	return title[:len(title)-18]
}