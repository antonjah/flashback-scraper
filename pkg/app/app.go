package app

import (
	"github.com/antonjah/flashback-scraper/pkg/config"
	"github.com/antonjah/flashback-scraper/pkg/output"
	"github.com/antonjah/flashback-scraper/pkg/thread"
)

func Run() {
	c := config.New()
	threads := thread.GetAll(c)
	output.Write(c, threads)
}
