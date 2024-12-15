package translator_test

import (
	"log/slog"
	"testing"

	"github.com/mmcdole/gofeed"
)

func TestGoFeed(t *testing.T) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://pythoncat.top/rss.xml")
	if err != nil {
		slog.Error("parse feed error", "error", err, "feed", feed)
	}
	slog.Info(feed.Title)
}
