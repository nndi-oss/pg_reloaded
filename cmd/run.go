package cmd

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func init() {
	runCmd.Flags().StringP("username", "u", "postgres", "Override the postgres user (default: postgres)")
	runCmd.Flags().StringP("host", "h", "localhost", "Override the server host (default: localhost)")
	runCmd.Flags().StringP("port", "p", "5432", "Override the server port (default: 5432)")

	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an immediate restore for a specific database",
	Long:  `Run an immediate restore for a specific database`,
	Run: func(cmd *cobra.Command, args []string) {
		dbName := args[0]
		var database DataseConfig
		for _, d := range config.Databases {
			if dbName == d.Name {
				database = d
				break
			}
		}
		if database == nil {
			fmt.Println("Invalid database specified. Run 'pg_reload list' to see configured databases")
			os.Exit(1)
			return
		}

		server = config.GetServerByName(database.Server)

		host := server.Host
		if 
		username := server.Username
		password := server.Password
		port := 5432
		sourceFile := database.Source.File

		err = RunDropDatabase(username, dbName, host, port, password)
		if err != nil {
			fmt.Println("Failed to drop database", err)
			os.Exit(1)
			return
		}
		err = RunPsql(username, dbName, host, port, sourceFile, password)
		if err != nil {
			fmt.Println("Failed to restore database", err)
			os.Exit(1)
			return
		}
	},
}
