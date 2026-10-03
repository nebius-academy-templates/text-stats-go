package main

import (
	"os"
	"strings"
)

// document carries text and logical lines through the report path.
type document struct {
	text  string
	lines []string
}

func documentFromText(text string) document {
	lines := make([]string, 0)
	for _, line := range splitLines(text) {
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}
	return document{text: text, lines: lines}
}

func documentFromPath(path string) (document, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return document{}, err
	}
	return documentFromText(strings.ToValidUTF8(string(raw), "")), nil
}

// sentenceFragments adapted from a Stack Overflow answer.
func (doc document) sentenceFragments() []string {
	return strings.Split(doc.text, ".")
}
