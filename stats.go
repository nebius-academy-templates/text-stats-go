package main

import (
	"strings"
	"unicode/utf8"
)

// Statistics functions are pure; only io.go reads files.
func countLines(lines []string) int { return len(lines) }

func countBlankLines(lines []string) int {
	count := 0
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			count++
		}
	}
	return count
}

func countCharacters(lines []string) int {
	count := 0
	for _, line := range lines {
		count += utf8.RuneCountInString(line)
	}
	return count
}

func countWords(lines []string) int {
	count := 0
	for _, line := range lines {
		count += len(strings.Fields(line))
	}
	return count
}
