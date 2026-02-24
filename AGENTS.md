# Agents Guide

This file contains guidance for AI agents (Claude, etc.) working on this project.
`CLAUDE.md` is a symlink to this file.

## Project Overview

**zettler** is a personal CLI tool for journaling and note-taking, built in Go.
It follows Zettelkasten principles and is designed to work alongside Obsidian.
The primary user workflow is daily journaling.

Module: `github.com/gormanity/zettler`

See `PROJECT.md` for architecture, design decisions, and planned commands.

## Development Workflow

### VCS: Jujutsu (`jj`)

- Use `jj` for all version control operations, not `git`.
- At the start of each work cycle, run `jj status` to check the working copy state.
- Only run `jj new` when the current change already has meaningful content — never
  create an empty change on top of another empty change.
- Commits should be **atomic and tightly scoped** — one logical change per commit.
- Use `jj describe -m "message"` to set the commit message, then `jj new` to open
  the next change.
- To push to main: `jj bookmark set main -r <rev>` then `jj git push --branch main`.
  Do NOT use `jj git push --change` — it creates unwanted push-* bookmarks.

### TDD

- Write a failing test first and commit it, then implement.
- **Do not push failing test commits** — the failing test commit and its paired
  implementation commit must both exist before pushing.
- Run tests with `make test` or `go test ./...`.

### Code Style

- Follow standard Go conventions (`gofmt`, `go vet`).
- Keep commands in `cmd/` as thin wrappers — business logic goes in internal packages.
- Prefer small, focused functions.
- Do not add comments unless the logic is non-obvious.

## Key Files

- `cmd/root.go` — Cobra root command and `Execute()`
- `cmd/*_test.go` — command-level tests
- `Makefile` — `make build`, `make test`, `make fmt`, `make install`
- `go.mod` — module `github.com/gormanity/zettler`, Go 1.23.4

## Running & Building

```sh
make test      # run all tests
make build     # build to bin/zettler
make install   # install to $GOPATH/bin
```
