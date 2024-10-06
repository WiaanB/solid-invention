package cmd

import (
	"cinnanym/database/surreal"
	"cinnanym/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the Web Server",
	Run: func(cmd *cobra.Command, args []string) {
		surreal.SetupDB()
		server.Serve()
	},
}
