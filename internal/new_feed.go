package main

import (
	"log/slog"

	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
)

type NewFeed struct {
}

func ParseFeedFromUrl(url string) {
	var newfeed *feeds.Feed

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)

	newfeed = &feeds.Feed{
		Title:       feed.Title,
		Link:        &feeds.Link{Href: feed.FeedLink},
		Description: feed.Description,
		Created:     *feed.PublishedParsed,
	}

	if len(feed.Authors) > 0 {
		author := feed.Authors[len(feed.Authors)-1]
		slog.Info("feed info", "title", author.Name, "email", author.Email)
		newfeed.Author = &feeds.Author{Name: author.Name, Email: author.Email}
	}

	for _, item := range feed.Items {
		newfeed.Add(&feeds.Item{
			Title:       item.Title,
			Link:        &feeds.Link{Href: item.Link},
			Description: item.Description,
			Created:     *item.PublishedParsed,
			// Categories:  []string{item.Categories[0]},
		})
	}

	rss, err := newfeed.ToRss()
	slog.Info("feed title", "xml-rss", rss, "err", err)
}
