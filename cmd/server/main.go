package main

import (
	"github.com/spf13/cobra"
	"github.com/xbmlz/gin-svelte-template/internal/bootstrap"
	"github.com/xbmlz/gin-svelte-template/internal/build"
	"github.com/xbmlz/gin-svelte-template/internal/core"
	"github.com/xbmlz/gin-svelte-template/internal/model"
	"go.uber.org/fx"
	"os"
)

var configPath string

func init() {
	rootCmd.Version = build.String()
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "C", "config.yaml", "config file path")

	for _, cmd := range []*cobra.Command{runCmd, migrateCmd} {
		rootCmd.AddCommand(cmd)
	}
}

var rootCmd = &cobra.Command{
	Short: "Gin Svelte Template Server",
	Long: `Gin Svelte Template Server
Usage:
  - 'server migrate' to run migrations
  - 'server run' to run the server
`,
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Long:  `Run the server`,
	Run: func(cmd *cobra.Command, args []string) {
		fx.New(bootstrap.Module).Run()
	},
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations",
	Long:  `Run migrations`,
	Run: func(cmd *cobra.Command, args []string) {
		config := core.NewConfig()
		log := core.NewLogger(config)
		db := core.NewDatabase(config, log)

		if err := db.DB.AutoMigrate(
			&model.User{},
		); err != nil {
			log.Errorf("failed to run migrations - %s", err)
		}
	},
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
