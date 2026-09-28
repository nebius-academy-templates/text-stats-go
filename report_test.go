package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCollectStatsKeepsDisplayOrder(t *testing.T) {
	values := collectStats([]string{"hello world", "", "again"})
	keys := make([]string, 0, len(values))
	counts := make([]int, 0, len(values))
	for _, value := range values {
		keys = append(keys, value.key)
		counts = append(counts, value.value)
	}
	if !reflect.DeepEqual(keys, []string{"lines", "blank_lines", "words", "characters"}) {
		t.Fatalf("keys = %v", keys)
	}
	if !reflect.DeepEqual(counts, []int{3, 1, 3, 16}) {
		t.Fatalf("counts = %v", counts)
	}
}

func TestFormatTextAlignsRows(t *testing.T) {
	rows := strings.Split(buildReport([]string{"one", "", "two"}), "\n")
	if len(rows) != 4 || !strings.HasPrefix(rows[0], "Lines") || !strings.HasSuffix(rows[0], "3") {
		t.Fatalf("report rows = %q", rows)
	}
	for _, row := range rows {
		if !strings.Contains(row, "  ") {
			t.Fatalf("row lacks aligned separator: %q", row)
		}
	}
}
