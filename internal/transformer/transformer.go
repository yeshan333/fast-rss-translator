package transformer

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/yeshan333/fast-rss-translator/internal/translator"
)

type Transformer struct {
	FilePath        string // handle filepath
	ModifyStartLine int
	ModifyEndLine   int
}

// func ListAllFeedFilesInDir(dirPath string) []string {
// 	fileList := []string{}

// 	entries, err := os.ReadDir(dirPath)
// 	if err != nil {
// 		slog.Error("Error reading directory", "path", dirPath, "err", err)
// 		return []string{}
// 	}

// 	for _, entry := range entries {
// 		if !entry.IsDir() {
// 			fileList = append(fileList, entry.Name())
// 		}
// 	}

// 	return fileList
// }

func ReplaceContentBetweenLines(filePath string, startLine, endLine int, newContent string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	lineNum := 1
	contentAdded := false

	for scanner.Scan() {
		line := scanner.Text()
		if lineNum < startLine || lineNum > endLine {
			lines = append(lines, line)
		} else if !contentAdded {
			// only replace content once
			lines = append(lines, newContent)
			contentAdded = true
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return WriteToFile(filePath, lines)
}

func WriteToFile(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}

func GetModifyRange(filePath string) (int, int) {
	var ModifyStartLine, ModifyEndLine int
	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("Error opening file", "path", filePath, "err", err)
		return 0, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "fast-rss-translator: start") {
			ModifyStartLine = lineNum
		}
		if strings.Contains(scanner.Text(), "fast-rss-translator: end") {
			ModifyEndLine = lineNum
		}
		if ModifyStartLine != 0 && ModifyEndLine != 0 {
			break
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		slog.Error("Error reading file", "path", filePath, "err", err)
		fmt.Println("Error reading file:", err)
	}

	return ModifyStartLine, ModifyEndLine
}

func DoTransform(modifyFilePath, visitBaseUrl string, feeds []translator.Feed) {
	visitInfo := ""

	for _, feed := range feeds {
		visitInfo = fmt.Sprintf("%s\n[%s](%s) -> [%s%s](%s%s)\n", visitInfo, feed.Url, feed.Url, visitBaseUrl, feed.Name, visitBaseUrl, feed.Name)
	}

	ModifyStartLine, ModifyEndLine := GetModifyRange(modifyFilePath)
	err := ReplaceContentBetweenLines(modifyFilePath, ModifyStartLine+1, ModifyEndLine-1, visitInfo)
	if err != nil {
		slog.Error("Error replacing content", "err", err)
	} else {
		slog.Info("Update file successfully", "path", modifyFilePath)
	}
}
