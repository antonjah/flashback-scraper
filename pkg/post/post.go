package post

import (
	"github.com/PuerkitoBio/goquery"
)

type Post struct {
	UserName  string `yaml:"user_name" json:"user_name"`
	TimeStamp string `yaml:"time_stamp" json:"time_stamp"`
	Content   string `yaml:"content" json:"content"`
}

func GetAll(doc *goquery.Document) []Post {
	var posts []Post

	doc.Find(".post").Each(func(_ int, s *goquery.Selection) {
		var p Post
		// Remove quoted texts
		s.Find(".post-bbcode-quote-wrapper").Remove()
		p.UserName = getUserName(s)
		p.TimeStamp = getTimeStamp(s)
		p.Content = getContent(s)

		posts = append(posts, p)
	})

	return posts
}
