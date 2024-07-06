package config

import "github.com/yeshan333/fast-rss-translator/internal/translator"

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
