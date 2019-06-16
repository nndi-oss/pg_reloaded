package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zikani03/pg_reloaded/pg_reloaded"
	"os"
)

func init() {
	// runCmd.Flags().StringP("username", "u", "postgres", "Override the postgres user (default: postgres)")
	// runCmd.Flags().StringP("host", "h", "localhost", "Override the server host (default: localhost)")
	// runCmd.Flags().StringP("port", "p", "5432", "Override the server port (default: 5432)")

	rootCmd.AddCommand(restoreCmd)
}

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restores a backup for a specific database without dropping first",
	Long:  `Restores a backup for a specific database without dropping first`,
	Run: func(cmd *cobra.Command, args []string) {
		dbName := args[0]
		var database pg_reloaded.DatabaseConfig
		var found bool = false
		for _, d := range config.Databases {
			if dbName == d.Name {
				database = d
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Invalid database specified. Run 'pg_reloaded list' to see configured databases")
			os.Exit(1)
			return
		}

		server := config.GetServerByName(database.Server)

		host := server.Host
		username := server.Username
		password := server.Password
		port := server.Port
		sourceFile := database.Source.File

		err := pg_reloaded.RunRestoreDatabase(username, dbName, host, port, sourceFile, password)
		if err != nil {
			fmt.Printf("Failed to restore database. Got %v", err)
			os.Exit(1)
			return
		}
	},
}
