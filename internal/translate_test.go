package main

import (
	"testing"

	translator "github.com/Conight/go-googletrans"
	"github.com/magiconair/properties/assert"
)

func TestGoogleTranslate(t *testing.T) {
	c := translator.Config{
		Proxy: "http://127.0.0.1:7890",
	}
	googleTranslator := translator.New(c)
	result, err := googleTranslator.Translate("你好 世界", "auto", "en")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, result.Text, "Hello World")
}
