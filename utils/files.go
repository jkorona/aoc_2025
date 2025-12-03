package utils

import (
	"fmt"
	"os"
	"bufio"
)

func ReadLinesFromFile(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []string{}
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return []string{}
	}
	return lines
}

func ReadFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return ""
	}
	return string(data)
}