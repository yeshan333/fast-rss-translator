# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based RSS feed translator that converts RSS feeds from one language to another. It supports multiple translation engines (Google Translate, Cloudflare AI, Alibaba Qwen) and can generate translated RSS feeds in various formats (RSS, Atom, JSON). The tool can be used as a standalone CLI application or as a GitHub Action.

## Architecture

The project follows a modular architecture with the following key components:

1. **Main Entry Point**: `main.go` initializes the application and calls the root command
2. **Command Structure**: Uses Cobra CLI framework (`cmd/` directory)
   - Root command handles the main translation workflow
   - Update command for updating files with translated feed URLs
3. **Configuration**: YAML-based configuration (`subscribes.yaml`) defining feeds to translate
4. **Core Components** (`internal/` directory):
   - `translator/`: Handles RSS feed parsing and translation logic
   - `config/`: Configuration structures and parsing
   - `transformer/`: Updates README files with translated feed URLs
5. **GitHub Action Integration**: Docker-based action with entrypoint script

## Common Development Tasks

### Building the Project

```bash
go build -o fast-rss-translator
```

### Running Tests

```bash
go test ./...
```

### Running the Application

```bash
# Basic usage
./fast-rss-translator --config subscribes.yaml --update-file README.md

# Using specific configuration
./fast-rss-translator --config ./path/to/config.yaml --update-file ./path/to/readme.md
```

### Docker Build

```bash
docker build -t fast-rss-translator .
```

## Key Files and Directories

- `main.go`: Application entry point
- `cmd/`: Cobra command structure
- `internal/translator/`: Core translation logic
- `internal/config/`: Configuration parsing
- `internal/transformer/`: README updating logic
- `subscribes.yaml`: Feed configuration
- `Dockerfile`: Containerization
- `action.yml`: GitHub Action definition
- `entrypoint.sh`: GitHub Action entrypoint script

## Translation Engines

The tool supports multiple translation engines:

1. Google Translate (default)
2. Cloudflare AI
3. Alibaba Qwen

Each engine has specific configuration requirements and API key handling.
