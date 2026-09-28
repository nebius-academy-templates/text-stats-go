package main

import (
	"fmt"
	"strings"
)

type statistic struct {
	key   string
	label string
	count func([]string) int
}

// A slice, not a map, makes report order an explicit contract.
var statistics = []statistic{
	{key: "lines", label: "Lines", count: countLines},
	{key: "blank_lines", label: "Blank lines", count: countBlankLines},
	{key: "words", label: "Words", count: countWords},
	{key: "characters", label: "Characters", count: countCharacters},
}

type statValue struct {
	key   string
	label string
	value int
}

func collectStats(lines []string) []statValue {
	values := make([]statValue, 0, len(statistics))
	for _, stat := range statistics {
		values = append(values, statValue{key: stat.key, label: stat.label, value: stat.count(lines)})
	}
	return values
}

func formatText(values []statValue) string {
	width := 0
	for _, item := range values {
		if len(item.label) > width {
			width = len(item.label)
		}
	}
	rows := make([]string, 0, len(values))
	for _, item := range values {
		rows = append(rows, fmt.Sprintf("%-*s  %d", width, item.label, item.value))
	}
	return strings.Join(rows, "\n")
}

func buildReport(lines []string) string { return formatText(collectStats(lines)) }
