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

- **softwaretestingweekly.xml**: [https://softwaretestingweekly.com/issues.rss](https://softwaretestingweekly.com/issues.rss) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/softwaretestingweekly.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/softwaretestingweekly.xml)

- **thinkingelixir.xml**: [https://podcast.thinkingelixir.com/rss](https://podcast.thinkingelixir.com/rss) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/thinkingelixir.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/thinkingelixir.xml)

- **netflixtechblog.xml**: [https://netflixtechblog.com/feed](https://netflixtechblog.com/feed) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/netflixtechblog.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/netflixtechblog.xml)

- **dzone.xml**: [https://feeds.dzone.com/home](https://feeds.dzone.com/home) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/dzone.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/dzone.xml)

- **microsoft_devblogs.xml**: [https://devblogs.microsoft.com/landingpage/](https://devblogs.microsoft.com/landingpage/) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/microsoft_devblogs.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/microsoft_devblogs.xml)

- **cncf_blog.xml**: [https://rsshub.rssforever.com/cncf](https://rsshub.rssforever.com/cncf) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/cncf_blog.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/cncf_blog.xml)

- **uber_blog.xml**: [https://rsshub.rssforever.com/uber/blog](https://rsshub.rssforever.com/uber/blog) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/uber_blog.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/uber_blog.xml)

- **grafana_blog.xml**: [https://grafana.com/blog/index.xml](https://grafana.com/blog/index.xml) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/grafana_blog.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/grafana_blog.xml)

- **ebpf_blog.xml**: [https://ebpf.io/blog/rss.xml](https://ebpf.io/blog/rss.xml) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/ebpf_blog.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/ebpf_blog.xml)

- **pythoncat.xml**: [https://pythoncat.top/rss.xml](https://pythoncat.top/rss.xml) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/pythoncat.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/pythoncat.xml)

- **systemdesign.xml**: [https://newsletter.systemdesign.one/feed](https://newsletter.systemdesign.one/feed) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/systemdesign.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/systemdesign.xml)

- **bytebytego.xml**: [https://blog.bytebytego.com/feed](https://blog.bytebytego.com/feed) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/bytebytego.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/bytebytego.xml)

- **sreweekly.xml**: [https://sreweekly.com/feed/](https://sreweekly.com/feed/) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/sreweekly.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/sreweekly.xml)

- **cloudflare.xml**: [https://blog.cloudflare.com/rss](https://blog.cloudflare.com/rss) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/cloudflare.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/cloudflare.xml)

- **javascriptweekly.xml**: [https://cprss.s3.amazonaws.com/javascriptweekly.com.xml](https://cprss.s3.amazonaws.com/javascriptweekly.com.xml) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/javascriptweekly.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/javascriptweekly.xml)

- **pycoders.xml**: [https://pycoders.com/feed/jDwcOsNM](https://pycoders.com/feed/jDwcOsNM) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/pycoders.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/pycoders.xml)

- **thisweekinreact.xml**: [https://thisweekinreact.com/newsletter/rss.xml](https://thisweekinreact.com/newsletter/rss.xml) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/thisweekinreact.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/thisweekinreact.xml)

- **elixirstatus.xml**: [https://elixirstatus.com/rss](https://elixirstatus.com/rss) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/elixirstatus.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/elixirstatus.xml)

- **sspai_daily.xml**: [https://rsshub.rssforever.com/sspai/author/ee0vj778](https://rsshub.rssforever.com/sspai/author/ee0vj778) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/sspai_daily.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/sspai_daily.xml)

- **golangweekly.xml**: [https://cprss.s3.amazonaws.com/golangweekly.com.xml](https://cprss.s3.amazonaws.com/golangweekly.com.xml) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/golangweekly.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/golangweekly.xml)

- **aws_architecture_blog.xml**: [https://aws.amazon.com/blogs/architecture/feed/](https://aws.amazon.com/blogs/architecture/feed/) -> [https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/aws_architecture_blog.xml](https://fastly.jsdelivr.net/gh/yeshan333/fast-rss-translator@main/rss/aws_architecture_blog.xml)

<!-- fast-rss-translator: end -->
```

more config introduction can be found in [action.yml](./action.yml).
