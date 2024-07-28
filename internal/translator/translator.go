package translator

import (
	"fmt"
	"log/slog"
	"math"
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
	TranslateMode   string `mapstructure:"translate_mode"`   // origin | proxy | bilingual, bilingual: mix origin and target lang, proxy: do not translate
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

	max := len(feed.Items)
	if translator.MaxPost < max {
		max = translator.MaxPost
	}

	wg := sync.WaitGroup{}
	wg.Add(max)
	// limit translate post items
	for i := 0; i < max; i++ {
		go func(i int) {
			var transTitle string
			var transDesc string
			if translator.TranslateMode == "bilingual" {
				transTitle = fmt.Sprintf("【%s】%s", feed.Items[i].Title, translator.DoTranslate(feed.Items[i].Title))
				transDesc = fmt.Sprintf("【%s】%s", feed.Items[i].Description, translator.DoTranslate(feed.Items[i].Description))
			} else if translator.TranslateMode == "proxy" {
				transTitle = feed.Items[i].Title
				transDesc = feed.Items[i].Description
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
				Proxy:       translator.HttpProxy,
				ServiceUrls: []string{"translate.google.com.hk"},
			}
			googleTranslator = gtranslator.New(c)
		} else {
			googleTranslator = gtranslator.New()
		}
		srcLang := "auto"
		if translator.OriginLanguage != "" {
			srcLang = translator.OriginLanguage
		}

		length := len(content)
		if length < 3000 {
			result, err := googleTranslator.Translate(content, srcLang, translator.TargetLanguage)
			if err != nil {
				slog.Error("use google translate err", "err", err, "feed", translator.Feed.Url, "translate_content", content)
				// return origin text
				return content
			}
			return result.Text
		} else {
			translatedContent := ""
			for i := 0; i < int(math.Ceil(float64(length)/3000.0)); i++ {
				start := i * 3000
				end := start + 3000
				if end > length {
					end = length
				}
				part := content[start:end]
				result, err := googleTranslator.Translate(part, srcLang, translator.TargetLanguage)
				if err != nil {
					slog.Error("use google translate err", "err", err, "feed", translator.Feed.Url, "translate_content", part)
					// return origin text
					return content
				}
				translatedContent += result.Text
			}
			return translatedContent
		}
	case "openai":
		return ""
	case "aliyun":
		return ""
	default:
		return ""
	}
}
