package app

import (
	"github.com/antonjah/flashback-scraper/pkg/config"
	"github.com/antonjah/flashback-scraper/pkg/thread"
)

func Run() {
	c := config.New()
	thread.GetAll(c)
}
