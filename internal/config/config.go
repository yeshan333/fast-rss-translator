package config

import "github.com/yeshan333/fast-rss-translator/internal/translator"

type Base struct {
	OutputPath string `mapstructure:"output_path"`
}

type Config struct {
	Base      Base               `mapstructure:"base"`
	HttpProxy string             `mapstructure:"http_proxy"`
	Feeds     []translator.Feeds `mapstructure:"feeds"`
}
