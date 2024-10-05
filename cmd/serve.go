package cmd

import (
	"cinnanym/database/postgres"
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
		postgres.InitialiseDB()
		server.Serve()
	},
}
