# Course requirements

Complete this setup once before starting the Codex Foundation lessons that use
`text-stats-go`.

## Go toolchain

This repository requires Go 1.26.8. The version is declared in `go.mod`:

```text
go 1.26
toolchain go1.26.8
```

While network access is available, install or cache that toolchain:

```sh
GOTOOLCHAIN=go1.26.8 go version
```

The command must report Go 1.26.8. Go may download the selected toolchain the
first time, so run this before a lesson or recording rather than during it.

Next, confirm that the repository works without downloading modules:

```sh
GOPROXY=off GOTOOLCHAIN=go1.26.8 go test ./...
```

If this command fails because the toolchain is unavailable, finish the Go
installation while network access is allowed, then repeat the offline check.

## No virtual environment or dependency install

The baseline application uses only the Go standard library. It does not need
`uv`, a Python virtual environment, or a Go-specific virtual environment. Do
not run `go get`, `go mod download`, or `go mod tidy` as course setup.

Lesson 5 deliberately contains an unverified synthetic dependency for a code
review exercise. Do not download, install, or execute that dependency. The
normal test and build paths do not need it.

## Codex project

Add the `text-stats-go` directory as its own project in the Codex app. Do not
reuse the Python `text-stats` project. Each `codex/lesson-N` branch is an
independent starting state, so do not carry code changes from one lesson into
another.

Lessons 3 and 4 include a prepared Codex environment that runs the offline test
command automatically. The other lessons assume the same one-time toolchain
preparation has already been completed.

## Standard lesson commands

Run the tests:

```sh
GOPROXY=off GOTOOLCHAIN=go1.26.8 go test ./...
```

Run the sample application:

```sh
GOPROXY=off GOTOOLCHAIN=go1.26.8 go run . samples/sample.txt
```

Build and run the binary when an explicit build is needed:

```sh
GOPROXY=off GOTOOLCHAIN=go1.26.8 go build .
./text-stats-go samples/sample.txt
```

The generated `text-stats-go` binary is ignored by Git.
