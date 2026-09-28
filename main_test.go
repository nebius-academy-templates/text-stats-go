package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunPrintsReport(t *testing.T) {
	path := filepath.Join(t.TempDir(), "notes.txt")
	if err := os.WriteFile(path, []byte("hello world\n\nagain\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	var stdout, stderr bytes.Buffer
	if code := run([]string{path}, &stdout, &stderr); code != 0 {
		t.Fatalf("exit code = %d, stderr = %q", code, stderr.String())
	}
	if stderr.Len() != 0 {
		t.Fatalf("unexpected stderr: %q", stderr.String())
	}
	for _, label := range []string{"Lines", "Blank lines", "Words", "Characters"} {
		if !strings.Contains(stdout.String(), label) {
			t.Fatalf("report lacks %q: %q", label, stdout.String())
		}
	}
}

func TestRunNeedsOnePath(t *testing.T) {
	for _, args := range [][]string{{}, {"one.txt", "two.txt"}} {
		var stdout, stderr bytes.Buffer
		if code := run(args, &stdout, &stderr); code != 2 {
			t.Fatalf("args %v: exit code = %d", args, code)
		}
		if !strings.Contains(stderr.String(), "usage:") {
			t.Fatalf("args %v: missing usage error", args)
		}
	}
}
