package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const maxFileBytes int64 = 5_000_000

// readText owns filesystem access and the input safeguards.
func readText(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if info.Size() > maxFileBytes {
		return "", fmt.Errorf("%s is %d bytes, over the %d byte limit", path, info.Size(), maxFileBytes)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	if int64(len(raw)) > maxFileBytes {
		return "", fmt.Errorf("%s is %d bytes, over the %d byte limit", path, len(raw), maxFileBytes)
	}
	if utf8.Valid(raw) {
		return string(raw), nil
	}

	var text strings.Builder
	text.Grow(len(raw))
	for _, b := range raw {
		text.WriteRune(rune(b))
	}
	return text.String(), nil
}

func readLines(path string) ([]string, error) {
	text, err := readText(path)
	if err != nil {
		return nil, err
	}
	return splitLines(text), nil
}

// splitLines follows the Python fixture's logical-line contract, including
// interior blank lines but not an extra line after a final line break.
func splitLines(text string) []string {
	if text == "" {
		return []string{}
	}
	lines := []string{}
	start := 0
	for i, r := range text {
		if i < start || !isLineBreak(r) {
			continue
		}
		lines = append(lines, text[start:i])
		start = i + utf8.RuneLen(r)
		if r == '\r' && start < len(text) && text[start] == '\n' {
			start++
		}
	}
	if start < len(text) {
		lines = append(lines, text[start:])
	}
	return lines
}

func isLineBreak(r rune) bool {
	switch r {
	case '\n', '\r', '\v', '\f', '\x1c', '\x1d', '\x1e', '\x85', '\u2028', '\u2029':
		return true
	default:
		return false
	}
}
