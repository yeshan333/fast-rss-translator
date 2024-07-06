FROM golang:1.22.4 as builder

WORKDIR /workspace
COPY . .
RUN go mod download
RUN go build -o fast-rss-translator

FROM alpine:3.20

LABEL "com.github.actions.name"="fast-rss-translator"
LABEL "com.github.actions.description"="fast-rss-translator"
LABEL "com.github.actions.icon"="home"
LABEL "com.github.actions.color"="green"

LABEL "repository"="https://github.com/yeshan333/fast-rss-translator"
LABEL "homepage"="https://github.com/yeshan333/fast-rss-translator"
LABEL "maintainer"="yeshan333.ye@gmail.com"

LABEL "Name"="fast-rss-translator"

ENV LC_ALL C.UTF-8
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US.UTF-8

RUN apk add --no-cache \
        git \
        openssh-client \
        libc6-compat \
        libstdc++ \
        curl

COPY --from=builder /workspace/fast-rss-translator /bin/fast-rss-translator
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh && chmod +x /bin/fast-rss-translator

ENTRYPOINT ["/entrypoint.sh"]