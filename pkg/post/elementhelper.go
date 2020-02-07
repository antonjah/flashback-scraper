package post

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const dateFormat = "2006-01-02"

func getUserName(doc *goquery.Selection) string {
	return doc.Find(".dropdown-header").Text()
}

func getTimeStamp(doc *goquery.Selection) string {
	d := strings.TrimSpace(strings.Split(doc.Find(".post-heading").Text(), "\n")[4])
	t := time.Now()
	isoT := t.Format(dateFormat)

	if strings.Contains(d, "Idag") {
		return strings.Replace(d[0:11], "Idag", isoT, -1)
	} else if strings.Contains(d, "Igår") {
		yd := t.AddDate(0, 0, -1)
		return strings.Replace(d[0:11], "Igår", yd.Format(dateFormat), -1)
	}

	return strings.Replace(d[0:17], ",", "", -1)
}

func getContent(doc *goquery.Selection) string {
	return strings.TrimSpace(doc.Find(".post_message").Text())
}
