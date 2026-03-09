# Project: zettler

A personal CLI tool for journaling and note-taking, built in Go.
Designed to support a Zettelkasten-style workflow and play nicely with Obsidian.

## Design Decisions

- Note format: plain Markdown with YAML frontmatter
- Filename convention: kebab-case slugs (`my-note-title.md`)
- ID scheme: slug is the identifier — no separate numeric/timestamp ID
- Editor: `$EDITOR` (or `editor` in config, falling back to `$EDITOR`)
- Config: `~/.config/zettler/config.toml`

## Config

```toml
vault = "/path/to/notes"
editor = ""  # optional, falls back to $EDITOR
```

## Vault Directory Structure

```
vault/
  inbox/       # default destination for `zettler new`
  journal/     # daily notes, named YYYY-MM-DD.md
  projects/    # PARA
  areas/       # PARA
  resources/   # PARA
  archive/     # PARA
```

## Note Format

```markdown
---
title: "Note Title"
created: 2026-02-24
tags: []
---
```

Journal entries use title format: `"Journal - February 24, 2026"`
They include frontmatter in this shape:

```markdown
---
title: "Journal - February 24, 2026"
created: 2026-02-24
location:
tags: []
---
```

## Commands

| Command | Behavior |
|---|---|
| `zettler new [title]` | Create note in `inbox/` (or `--dir` for relative path from vault root, `.` for cwd), open in `$EDITOR` |
| `zettler journal [date\|yesterday]` | Open journal entry for today (or given date), create if missing |
| `zettler edit <slug>` | Open note by exact filename slug (without `.md`) in `$EDITOR` |
| `zettler list` | List notes in vault |
| `zettler scratch` | Open a throwaway file in `$TMPDIR` in `$EDITOR` |
