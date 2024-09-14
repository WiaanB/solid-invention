package cmd

import (
	"github.com/spf13/cobra"
	"gotcha/database/postgres"
	"gotcha/server"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves the Gotcha Web Server",
	Run: func(cmd *cobra.Command, args []string) {
		postgres.InitialiseDB()
		server.Serve()
	},
}
