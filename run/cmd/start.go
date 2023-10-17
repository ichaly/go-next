package cmd

import (
	"github.com/ichaly/go-next/app"
	"github.com/ichaly/go-next/pkg"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"path/filepath"
)

var configFile string

var runCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Service.",

	Run: func(cmd *cobra.Command, args []string) {
		if configFile == "" {
			configFile = filepath.Join("./cfg", "dev.yml")
		}
		fx.New(
			app.Modules,
			pkg.Modules,
			fx.Supply(configFile),
		).Run()
	},
}

func init() {
	runCmd.PersistentFlags().StringVarP(
		&configFile, "config", "c", "", "start app with config file",
	)
	rootCmd.AddCommand(runCmd)
}
