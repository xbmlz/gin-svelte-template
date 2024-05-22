package main

import (
	"github.com/xbmlz/gin-svelte-template/internal/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		bootstrap.Module,
		fx.NopLogger,
	).Run()
}
