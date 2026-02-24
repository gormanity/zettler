# zettler

A personal CLI for journaling and note-taking in the Zettelkasten style. Notes are plain Markdown with YAML frontmatter, stored in a vault directory of your choosing. Designed to work alongside Obsidian.

## Installation

```sh
git clone https://github.com/gormanity/zettler.git
cd zettler
make install
```

Requires Go 1.20+. The binary is installed to `$GOPATH/bin`.

## Configuration

Create `~/.config/zettler/config.toml`:

```toml
vault = "/path/to/your/notes"
editor = ""  # optional, falls back to $EDITOR
```

The vault can be any directory. Zettler will create `inbox/` and `journal/` subdirectories as needed.

## Commands

| Command | Description |
|---|---|
| `zettler new [title]` | Create a note in `inbox/`. Use `--dir` for a path relative to vault root, or `.` for the current directory. |
| `zettler journal [date]` | Open today's journal entry, creating it if missing. Accepts `yesterday` or `YYYY-MM-DD`. |
| `zettler edit <slug>` | Find and open a note by its filename slug. |
| `zettler list` | List all notes in the vault. |
| `zettler scratch` | Open a throwaway file in `$TMPDIR`. |

## Note Format

```markdown
---
title: "My Note"
created: 2026-02-25
tags: []
---
```

Filenames use kebab-case slugs derived from the title (e.g. `my-note.md`).
