/*
Copyright © 2024 yeshan333 <yeshan333.ye@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yeshan333/fast-rss-translator/cmd/commands"
	"github.com/yeshan333/fast-rss-translator/internal/config"
	"github.com/yeshan333/fast-rss-translator/internal/transformer"
	"github.com/yeshan333/fast-rss-translator/internal/translator"
)

var cfgFile string // rss feed subcribes configuration file
var globalConfig config.Config
var updateFile string

var rootCmd = &cobra.Command{
	Use:   "fast-rss-translator",
	Short: "A faster rss translator for rss feeds. translate any xml-format file",
	Run: func(cmd *cobra.Command, args []string) {
		// cfgFile, err := cmd.Flags().GetString("config")
		// if err != nil {
		// 	slog.Error("somthing wrong", "err", err)
		// 	panic(err)
		// }
		slog.Info("get rss feeds from config file", "filepath", cfgFile)
		slog.Info("global subscribes config", "config", fmt.Sprintf("%+v", globalConfig))

		wg := sync.WaitGroup{}
		wg.Add(len(globalConfig.Feeds))
		for i := 0; i < len(globalConfig.Feeds); i++ {
			// default translate engine
			translate_engine := globalConfig.Base.TranslateEngine

			go func(i int) {
				if globalConfig.Feeds[i].TranslateEngine == "" {
					globalConfig.Feeds[i].TranslateEngine = translate_engine
				}
				trans := &translator.Translator{
					Feed:      globalConfig.Feeds[i],
					HttpProxy: globalConfig.HttpProxy,
				}
				trans.Execute(globalConfig.Base.OutputPath)
				wg.Done()
			}(i)
		}
		wg.Wait()

		transformer.DoTransform(
			updateFile,
			globalConfig.Base.VisitBasicUrl,
			globalConfig.Feeds,
		)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func ReadConfig() {
	viper.SetConfigFile(cfgFile)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("somthing wrong", "err", err)
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&globalConfig)
	if err != nil {
		slog.Error("somthing wrong", "err", err)
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func init() {
	rootCmd.AddCommand(commands.UpdateCmd)
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "subscribes.yaml", "config file (default is $(pwd)/subscribes.yaml)")
	rootCmd.Flags().StringVarP(&updateFile, "update-file", "f", "README.md", "update file path (default is $(pwd)/README.md)")
	ReadConfig()
	rootCmd.MarkFlagRequired("update-file")
}
