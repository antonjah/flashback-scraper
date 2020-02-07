package post

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Post struct {
	UserName  string `json:"user_name"`
	TimeStamp string `json:"time_stamp"`
	Content   string `json:"content"`
}

func (p *Post) ToString() string {
	return fmt.Sprintf("%s | %s\n%s\n", p.UserName, p.TimeStamp, p.Content)
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
