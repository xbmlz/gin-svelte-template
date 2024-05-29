FROM golang:1.19-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /src

ADD . /src

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

EXPOSE 9080

ENTRYPOINT ["/entrypoint.sh"]