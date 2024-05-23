package main

import (
	"github.com/spf13/cobra"
	"github.com/xbmlz/gin-svelte-template/build"
	"github.com/xbmlz/gin-svelte-template/internal/bootstrap"
	"go.uber.org/fx"
)

var configPath string

func init() {
	rootCmd.Version = build.String()
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "C", "config.yaml", "config file path")

	for _, cmd := range []*cobra.Command{runCmd} {
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
		fx.New(bootstrap.Module, fx.NopLogger).Run()
	},
}
