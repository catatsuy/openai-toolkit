package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Message represents a single message entry for JSONL
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// JSONLData represents the structure for JSONL format
type JSONLData struct {
	Messages []Message `json:"messages"`
}

// readFileContent reads the file and returns its content as a string
func readFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var contentBuilder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contentBuilder.WriteString(scanner.Text())
		contentBuilder.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	// Remove the trailing newline if it exists
	content := contentBuilder.String()
	content = strings.TrimSuffix(content, "\n")

	return content, nil
}

// generateJSONL generates the JSONL formatted output from file content
func generateJSONL(fileContent string) (string, error) {
	data := JSONLData{
		Messages: []Message{
			{Role: "user", Content: fileContent},
		},
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <file1> <file2> ...")
		return
	}

	files := os.Args[1:]

	for _, file := range files {
		content, err := readFileContent(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", file, err)
			continue
		}

		jsonlContent, err := generateJSONL(content)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating JSONL for %s: %v\n", file, err)
			continue
		}

		// Print JSONL to standard output
		fmt.Println(jsonlContent)
	}
}
