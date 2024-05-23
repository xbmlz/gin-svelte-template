package main

import (
	"os"
)

// @title           Gin Svelte Template API
// @version         1.0.0
// @description     This is a sample server.
// @termsOfService  http://swagger.io/terms/
// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io
// @license.name    MIT
// @license.url     http://www.apache.org/licenses/MIT.html
// @host            localhost:8765
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
