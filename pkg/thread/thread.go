package thread

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"

	"github.com/antonjah/flashback-scraper/pkg/config"
	"github.com/antonjah/flashback-scraper/pkg/post"
)

type Thread struct {
	Title string      `yaml:"title" json:"title"`
	Posts []post.Post `yaml:"posts" json:"posts"`
}

type result struct {
	thread Thread
	error  error
}

// GetAll loads all URLs from the config
func GetAll(c config.Config) []Thread {
	var rChan = make(chan result)
	var threads []Thread

	for _, url := range c.URLs {
		go parse(url, rChan)
		if r := <-rChan; r.error == nil {
			threads = append(threads, r.thread)
		} else {
			log.Error(r.error)
		}
	}
	close(rChan)

	return threads
}

func parse(url string, rChan chan result) {
	r, err := http.Get(url)
	if err != nil {
		rChan <- result{Thread{}, err}
		return
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		rChan <- result{Thread{}, errors.New(url + " - " + r.Status)}
		return
	}

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		rChan <- result{Thread{}, err}
		return
	}

	var t Thread
	t.Title = getTitle(doc)
	t.Posts = post.GetAll(doc)

	rChan <- result{t, nil}
}

func getPageCount() {}
