package main

import (
	"os"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
