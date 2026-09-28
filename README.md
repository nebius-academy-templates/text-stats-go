# text-stats-go

A small, standard-library-only command-line tool for reporting statistics about a text file. This is the local Go fixture for the Codex Foundation lessons; it is independent of the Python `text-stats` repository.

## Before the lessons

Complete the one-time setup in [Course requirements](REQUIREMENTS.md). It
prepares Go 1.26.8, verifies the repository offline, and explains why this
project does not need `uv`, a virtual environment, or a dependency-install
step.

## Run and test

```sh
go run . samples/sample.txt
go test ./...
go vet ./...
```

The four report rows appear in this order: Lines, Blank lines, Words, Characters. Blank lines count as lines. Character counts exclude line breaks and count Unicode code points, not UTF-8 bytes. Input is capped at five million bytes; invalid UTF-8 input falls back to Latin-1.

Each `codex/lesson-N` branch is an independent starting state. Do not carry changes from one lesson into another. The local `example.com` module path is a training placeholder, not a publishing destination.
