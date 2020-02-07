package post

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getUserName(doc *goquery.Selection) string {
	return doc.Find(".dropdown-header").Text()
}

func getTimeStamp(doc *goquery.Selection) string {
	return strings.TrimSpace(strings.Split(doc.Find(".post-heading").Text(), "\n")[4])
}

func getContent(doc *goquery.Selection) string {
	return strings.TrimSpace(doc.Find(".post_message").Text())
}
