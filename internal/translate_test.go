package main

import (
	"log/slog"
	"testing"

	translator "github.com/Conight/go-googletrans"
	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/assert"
)

func TestParseFeed(t *testing.T) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://shansan.top/rss2.xml")

	if err != nil {
		slog.Error("parse feed raise exception", "err", err, "feedUrl", "https://pythoncat.top/rss.xml")
		return
	}
	t.Log(feed.Description)
}

func TestGoogleTranslate(t *testing.T) {
	c := translator.Config{
		Proxy: "http://127.0.0.1:7890",
	}
	googleTranslator := translator.New(c)
	result, err := googleTranslator.Translate("hello", "auto", "zh")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, result.Text, "你好")
}
