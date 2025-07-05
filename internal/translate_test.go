package main

import (
	"log/slog"
	"os" // Added import for os
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
	if os.Getenv("CI") != "" || os.Getenv("HTTP_PROXY") == "" {
		t.Skip("Skipping Google Translate integration test in CI or when HTTP_PROXY is not set")
	}
	c := translator.Config{
		Proxy: os.Getenv("HTTP_PROXY"),
	}
	googleTranslator := translator.New(c)
	result, err := googleTranslator.Translate("hello", "auto", "zh")
	if err != nil {
		// Instead of panic, log the error and fail the test
		t.Fatalf("Google Translate request failed: %v", err)
	}

	assert.Equal(t, "你好", result.Text) // Corrected assertion order
}
