package cmd

import (
	"github.com/kolobok-kelbek/go-example-service/di"
	"github.com/spf13/cobra"
	"os"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "up server",
	Run: func(cmd *cobra.Command, args []string) {
		app := di.Init()
		err := app.Run()
		if err != nil {
			os.Exit(1)
		}
	},
}
