package translator

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
)

type Feeds struct {
	Name            string `mapstructure:"name"`
	Url             string `mapstructure:"url"`
	TargetLanguage  string `mapstructure:"target_language"`
	TranslateMode   string `mapstructure:"translate_mode"`   // origin | bilingual, mix origin and target lang
	TranslateEngine string `mapstructure:"translate_engine"` // google | openai
	MaxPost         int    `mapstructure:"max_post"`         // max handled posts
}

type Translator struct {
	Feeds
}

func (translator *Translator) Execute(outputDir string) {
	var newfeed *feeds.Feed

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(translator.Url)

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

	// limit translate post items
	for i := 0; i < translator.MaxPost; i++ {
		newfeed.Add(&feeds.Item{
			Title:       feed.Items[i].Title,
			Link:        &feeds.Link{Href: feed.Items[i].Link},
			Description: feed.Items[i].Description,
			Created:     *feed.Items[i].PublishedParsed,
			// Categories:  []string{item.Categories[0]},
		})
	}

	// var rss string
	// switch feed.FeedType {
	// case gofeed.FeedTypeAtom:
	// 	rss, _ := newfeed.ToRss()
	// 	CreateNewFeedFile(rss, targetFile)
	// default:
	// 	rss, _ := newfeed.ToRss()
	// 	CreateNewFeedFile(rss, targetFile)
	// }
	rss, err := newfeed.ToRss()
	if err != nil {
		slog.Error("parse rss raise exception", "err", err)
	}

	translator.CreateNewFeedFile(rss, fmt.Sprintf("%s%c%s", outputDir, filepath.Separator, translator.Name))
}

// targetFile: use absolute path
func (translator *Translator) CreateNewFeedFile(rssContent, targetFile string) error {
	fileName := targetFile

	dirPath := filepath.Dir(fileName)

	// if dir no exist, create it~
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	if err := os.WriteFile(fileName, []byte(rssContent), 0644); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (translator *Translator) DoTranslate() string {
	switch translator.TranslateEngine {
	case "Google":
		return ""
	case "OpenAI":
		return ""
	default:
		return ""
	}
}
