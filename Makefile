
run:
	go mod tidy
	go mod download
	GIN_MODE=debug CGO_ENABLE=0 go run ./cmd/server

.PHONY: run