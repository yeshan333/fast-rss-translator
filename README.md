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

**config_file**: A YAML file path that defines the RSS feeds to be translated and the translation settings. See the example: [subscribes.yaml](./subscribes.yaml).
This file contains global settings under the `base` key and a list of feed-specific configurations under the `feeds` key.

Supported translation engines (`translate_engine`):

- `google`: Uses Google Translate. Requires `http_proxy` to be set if accessing from certain regions.
- `cloudflare`: Uses Cloudflare Worker AI for translation.
  - Requires `cloudflare_account_id` and `cloudflare_api_key` to be set either globally under `base` or for a specific feed.
  - Alternatively, these can be set as environment variables: `CLOUDFLARE_ACCOUNT_ID` and `CLOUDFLARE_API_KEY`. Feed-specific configurations take precedence.

For detailed structure of the `subscribes.yaml` file and all available options, please refer to the [subscribes.yaml](./subscribes.yaml) example.

Make sure **update_file** has content section as follow, example: [TEST_README.md](./TEST_README.md):

```text
<!-- fast-rss-translator: start -->
new translated feed subscribe url will write here.
<!-- fast-rss-translator: end -->
```

more config introduction can be found in [action.yml](./action.yml).
