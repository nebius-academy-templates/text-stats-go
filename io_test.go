package main

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func writeInput(t *testing.T, name string, contents []byte) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, contents, 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestReadTextUTF8(t *testing.T) {
	path := writeInput(t, "utf8.txt", []byte("café\nnaïve\n"))
	text, err := readText(path)
	if err != nil || text != "café\nnaïve\n" {
		t.Fatalf("readText = %q, %v", text, err)
	}
}

func TestReadTextLatin1Fallback(t *testing.T) {
	path := writeInput(t, "latin1.txt", []byte{'c', 'a', 'f', 0xe9})
	text, err := readText(path)
	if err != nil || text != "café" {
		t.Fatalf("readText = %q, %v", text, err)
	}
}

func TestReadTextRejectsOversize(t *testing.T) {
	path := writeInput(t, "large.txt", []byte(strings.Repeat("x", int(maxFileBytes)+1)))
	if _, err := readText(path); err == nil || !strings.Contains(err.Error(), "byte limit") {
		t.Fatalf("expected byte-limit error, got %v", err)
	}
}

func TestReadLinesPreservesInteriorBlank(t *testing.T) {
	path := writeInput(t, "lines.txt", []byte("one\n\nthree\n"))
	lines, err := readLines(path)
	if err != nil || !reflect.DeepEqual(lines, []string{"one", "", "three"}) {
		t.Fatalf("readLines = %q, %v", lines, err)
	}
}

func TestReadDocumentOrdinaryInput(t *testing.T) {
	path := writeInput(t, "document.txt", []byte("one\ntwo\n"))
	doc, err := readDocument(path)
	if err != nil || !reflect.DeepEqual(doc.lines, []string{"one", "two"}) {
		t.Fatalf("readDocument = %#v, %v", doc, err)
	}
}

func TestSplitLinesCRLF(t *testing.T) {
	got := splitLines("one\r\n\r\nthree\r\n")
	if !reflect.DeepEqual(got, []string{"one", "", "three"}) {
		t.Fatalf("splitLines = %q", got)
	}
}

func TestSplitLinesEmpty(t *testing.T) {
	if got := splitLines(""); len(got) != 0 {
		t.Fatalf("splitLines = %q", got)
	}
}

func TestSplitLinesUnicodeSeparator(t *testing.T) {
	got := splitLines("one\u2028two")
	if !reflect.DeepEqual(got, []string{"one", "two"}) {
		t.Fatalf("splitLines = %q", got)
	}
}
