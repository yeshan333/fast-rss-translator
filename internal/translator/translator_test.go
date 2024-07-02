package translator_test

import (
	"testing"

	"github.com/yeshan333/fast-rss-translator/internal/translator"
)

func TestExecute(t *testing.T) {
	trans := &translator.Translator{
		Feed: translator.Feed{
			Name:            "feed_test.xml",
			Url:             "https://shan333.cn/rss2.xml",
			TargetLanguage:  "en",
			TranslateMode:   "origin",
			TranslateEngine: "google",
			MaxPost:         1,
		},
	}
	trans.Execute("./rss_test")
}
