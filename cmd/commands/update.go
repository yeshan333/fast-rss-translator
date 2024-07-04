package commands

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var filepath string
var cdn string
var FeedsFileLocateDir string

var UpdateCmd = &cobra.Command{
	Use:   "update-file",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		updateFilePath, err := cmd.Flags().GetString("filepath")
		if err != nil {
			slog.Error("somthing wrong", "err", err)
			panic(err)
		}

		slog.Info("updateFilePath", "path", updateFilePath)
	},
}

func init() {
	UpdateCmd.Flags().StringVarP(&filepath, "filepath", "f", "README.md", "update file path (default is $(pwd)/README.md)")
	UpdateCmd.Flags().StringVarP(&cdn, "cdn", "n", "jsdelive", "use cdn link to replace github link")
	UpdateCmd.Flags().StringVarP(&FeedsFileLocateDir, "dir", "d", "./rss", "feed xml file dir")
}
