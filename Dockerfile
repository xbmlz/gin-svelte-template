FROM golang:1.19-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /src

ADD . /src

RUN apk --no-cache add build-base git bash nodejs npm && npm install -g pnpm@8.9.2

RUN make build


FROM alpine:latest

ENV TIMEZONE="Asia/Shanghai"

RUN apk update \
    && apk --no-cache add \
        bash \
        ca-certificates \
        curl \
        tzdata \
    && ln -sf /usr/share/zoneinfo/${TIMEZONE} /etc/localtime \
    && echo "${TIMEZONE}" > /etc/timezone


COPY --from=builder /src/server /app
COPY --from=builder /src/config.yaml /app
COPY --from=builder /src/script/entrypoint.sh /app

WORKDIR /app

EXPOSE 8765

ENTRYPOINT ["/entrypoint.sh"]