package main

import "testing"

func TestCountLinesAndBlankLines(t *testing.T) {
	lines := []string{"one", "", " \t", "three"}
	if countLines(lines) != 4 || countBlankLines(lines) != 2 {
		t.Fatalf("counts = %d lines, %d blanks", countLines(lines), countBlankLines(lines))
	}
}

func TestCountCharactersUsesRunes(t *testing.T) {
	if got := countCharacters([]string{"café", "😊"}); got != 5 {
		t.Fatalf("character count = %d", got)
	}
}

func TestCountWordsPreservesPunctuation(t *testing.T) {
	if got := countWords([]string{"don't stop, he said."}); got != 4 {
		t.Fatalf("word count = %d", got)
	}
}

func TestEmptyStats(t *testing.T) {
	if countLines(nil) != 0 || countBlankLines(nil) != 0 || countCharacters(nil) != 0 || countWords(nil) != 0 {
		t.Fatal("empty input must have zero counts")
	}
}
