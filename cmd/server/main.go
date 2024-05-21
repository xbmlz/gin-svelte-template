package main

import (
	"github.com/xbmlz/gin-svelte-template/pkg/logger"
)

func main() {
	// fx.New(
	// 	bootstrap.Modules,
	// ).Run()
	log := logger.NewLogger(
		logger.DebugLevel,
	)

	log.Info("Server started")
}
