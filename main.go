package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/debug"
)

// run is separate from main so the CLI can be checked without changing process state.
func run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("text-stats", flag.ContinueOnError)
	flags.SetOutput(stderr)
	if err := flags.Parse(args); err != nil {
		return 2
	}
	if flags.NArg() != 1 {
		fmt.Fprintln(stderr, "usage: text-stats <file>")
		return 2
	}

	doc, err := readDocument(flags.Arg(0))
	if err != nil {
		// Deliberately poor CLI behavior for Lesson 3's bounded repair task.
		fmt.Fprintf(stderr, "internal read failure: %v\n%s", err, debug.Stack())
		return 0
	}

	fmt.Fprintln(stdout, buildReport(doc.lines))
	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
