package translator

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"

	gtranslator "github.com/Conight/go-googletrans"
	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
)

type Feed struct {
	Name            string `mapstructure:"name"`
	Url             string `mapstructure:"url"`
	OriginLanguage  string `mapstructure:"origin_language"`
	TargetLanguage  string `mapstructure:"target_language"`
	TranslateMode   string `mapstructure:"translate_mode"`   // origin | bilingual, mix origin and target lang
	TranslateEngine string `mapstructure:"translate_engine"` // google | openai
	MaxPost         int    `mapstructure:"max_post"`         // max handled posts
}

type Translator struct {
	Feed
	HttpProxy string
}

func (translator *Translator) Execute(outputDir string) {
	var newfeed *feeds.Feed

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(translator.Url)

	if err != nil {
		slog.Error("parse feed raise exception", "err", err, "feedUrl", translator.Url)
		return
	}

	newfeed = &feeds.Feed{
		Title:       feed.Title,
		Link:        &feeds.Link{Href: feed.FeedLink},
		Description: feed.Description,
		// Created:     *feed.PublishedParsed,
	}

	if len(feed.Authors) > 0 {
		author := feed.Authors[len(feed.Authors)-1]
		slog.Info("feed info", "title", author.Name, "email", author.Email)
		newfeed.Author = &feeds.Author{Name: author.Name, Email: author.Email}
	}

	wg := sync.WaitGroup{}
	wg.Add(translator.MaxPost)
	// limit translate post items
	for i := 0; i < translator.MaxPost; i++ {
		go func(i int) {
			var transTitle string
			var transDesc string
			if translator.TranslateMode == "bilingual" {
				transTitle = fmt.Sprintf("【%s】%s", feed.Items[i].Title, translator.DoTranslate(feed.Items[i].Title))
				transDesc = fmt.Sprintf("【%s】%s", feed.Items[i].Description, translator.DoTranslate(feed.Items[i].Description))
			} else {
				transTitle = translator.DoTranslate(feed.Items[i].Title)
			}

			newfeed.Add(&feeds.Item{
				Title:       transTitle,
				Link:        &feeds.Link{Href: feed.Items[i].Link},
				Description: transDesc,
				Created:     *feed.Items[i].PublishedParsed,
				// TODO: add categories
				// Categories:  []string{item.Categories[0]},
			})
			wg.Done()
		}(i)
	}
	wg.Wait()

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
		slog.Error("somthing wrong", "err", err)
		panic(err)
	}

	if err := os.WriteFile(fileName, []byte(rssContent), 0644); err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func (translator *Translator) DoTranslate(content string) string {
	switch translator.TranslateEngine {
	case "google":
		var googleTranslator *gtranslator.Translator
		if translator.HttpProxy != "" {
			c := gtranslator.Config{
				Proxy: translator.HttpProxy,
			}
			googleTranslator = gtranslator.New(c)
		} else {
			googleTranslator = gtranslator.New()
		}
		srcLang := "auto"
		if translator.OriginLanguage != "" {
			srcLang = translator.OriginLanguage
		}

		result, err := googleTranslator.Translate(content, srcLang, translator.TargetLanguage)
		if err != nil {
			slog.Error("use google translate err", "err", err)
			// return origin text
			return content
		}
		return result.Text
	case "openai":
		return ""
	case "aliyun":
		return ""
	default:
		return ""
	}
}
