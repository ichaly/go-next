package cli

import (
	"github.com/ichaly/go-next/api"
	"github.com/ichaly/go-next/lib"
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
			api.Modules,
			lib.Modules,
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
