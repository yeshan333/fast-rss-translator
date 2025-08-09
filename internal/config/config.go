package config

import (
	"fmt"
	"strings"

	"github.com/yeshan333/fast-rss-translator/internal/translator"
)

type Base struct {
	OutputPath      string `mapstructure:"output_path"`
	VisitBasicUrl   string `mapstructure:"visit_base_url"`
	TranslateEngine string `mapstructure:"translate_engine"`
}

type Config struct {
	Base      Base              `mapstructure:"base"`
	HttpProxy string            `mapstructure:"http_proxy"`
	Feeds     []translator.Feed `mapstructure:"feeds"`
}

// SafeString returns a string representation of the config with sensitive data masked
func (c *Config) SafeString() string {
	var maskedFeeds []string
	for _, feed := range c.Feeds {
		maskedFeeds = append(maskedFeeds, feed.SafeString())
	}
	
	return fmt.Sprintf("Config{Base: %+v, HttpProxy: %s, Feeds: [%s]}", 
		c.Base, c.HttpProxy, strings.Join(maskedFeeds, ", "))
}
