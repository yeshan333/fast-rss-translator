package transformer_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeshan333/fast-rss-translator/internal/transformer"
)

func TestGetModifyRange(t *testing.T) {
	ModifyStartLine, ModifyEndLine := transformer.GetModifyRange("/home/yeshan333/workspace/github/fast-rss-translator/README.md")
	assert.Equal(t, ModifyStartLine, 0)
	assert.Equal(t, ModifyEndLine, 0)
}

func TestReplaceContent(t *testing.T) {
	filePath := "/home/yeshan333/workspace/github/fast-rss-translator/README.md" // 替换为你的文件路径
	ModifyStartLine, ModifyEndLine := transformer.GetModifyRange(filePath)
	newContent := "nanding" // 新内容

	err := transformer.ReplaceContentBetweenLines(filePath, ModifyStartLine+1, ModifyEndLine-1, newContent)
	if err != nil {
		fmt.Println("Error replacing content:", err)
	} else {
		fmt.Println("Content replaced successfully.")
	}
}

// func TestListAllFeedFilesInDir(t *testing.T) {
// 	workspace, _ := os.Getwd()
// 	feedsLocateDir := workspace
// 	t.Log(transformer.ListAllFeedFilesInDir(feedsLocateDir))
// }
