package main

import (
	"testing"

	translator "github.com/Conight/go-googletrans"
	"github.com/stretchr/testify/assert"
)

func TestGoogleTranslate(t *testing.T) {
	c := translator.Config{
		Proxy: "http://127.0.0.1:7890",
	}
	googleTranslator := translator.New(c)
	result, err := googleTranslator.Translate("hello", "auto", "zh")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, result.Text, "你好")
}
