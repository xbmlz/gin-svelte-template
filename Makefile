.PHONY: install run build swag

VERSION=1.0.0
BIN=./bin/server
APP_MAIN=./cmd/server
GO_MODULE=github.com/xbmlz/gin-svelte-template

GO_ENV=CGO_ENABLED=0 GO111MODULE=on
Revision=$(shell git rev-parse --short HEAD 2>/dev/null || echo "")
GO_FLAGS=-ldflags="-X $(GO_MODULE)/build.Version=$(VERSION) -X '$(GO_MODULE)/build.Revision=$(Revision)' -X '$(GO_MODULE)/build.Time=`date +%FT%T%z`' -extldflags -static"
GO=$(GO_ENV) go

install:
	@$(GO) install github.com/swaggo/swag/cmd/swag@latest

run:swag
	@$(GO) run $(APP_MAIN) run

build:
	@$(GO) build $(GO_FLAGS) -o $(BIN) $(MAIN_PKG)

swag:
	@swag init -g $(APP_MAIN)/main.go -o docs

migrate:
	@$(GO) run $(APP_MAIN) migrate