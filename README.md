# Faster-RSS-Translator

A faster RSS translator for translating any language feed to any language feed with GitHub Action Automation Workflow.

Support Feed format:
- RSS (0.90 to 2.0)
- Atom (0.3, 1.0)
- JSON (1.0, 1.1)

## Usage

Here is a example github action configuration for you.

```yaml
name: test_github_action

on:
  push:
    branches: [ main ]
  workflow_dispatch:

permissions:
  contents: write

jobs:
  fast_rss_translator_action:
    runs-on: ubuntu-latest
    name: test fast-rss-translator action
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: fast-rss-translator
        uses: yeshan333/fast-rss-translator@main
        id: fast-rss-translator-action
        with:
          config_file: 'subscribes.yaml'
          update_file: 'README.md'
          push: true
          username: "github-actions[bot]"
          org: "yeshan333"
          repo: "fast-rss-translator"
          token: ${{ secrets.GITHUB_TOKEN }}
```

**config_file**: a file path. for reading origin feed url and translate it. example: [subscribes.yaml](./subscribes.yaml)

Make sure **update_file** has content section as follow, example: [TEST_README.md](./TEST_README.md):

```text
<!-- fast-rss-translator: start -->
new translated feed subscribe url will write here.
<!-- fast-rss-translator: end -->
```

more config introduction can be found in [action.yml](./action.yml).
