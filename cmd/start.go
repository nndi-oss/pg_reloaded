package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zikani03/pg_reloaded/pg_reloaded"
	"github.com/zikani03/pg_reloaded/cron"
)

var scheduler = cron.New()

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the PG Reloaded scheduler/cron daemon",
	Long:  `Starts the PG Reloaded scheduler/cron daemon`,
	Run: func(cmd *cobra.Command, args []string) {
		createJobsFromConfiguration(scheduler)
		scheduler.Start()
	},
}

func createJobsFromConfiguration(cronScheduler *cron.Cron) error {
	var server pg_reloaded.ServerConfig
	for _, db := range config.Databases {
		server = *config.GetServerByName(db.Server)
		username := server.Username
		password := server.Password
		err := cronScheduler.AddFunc(db.Schedule, func() {
			if db.Source.Type == "sql" {
				pg_reloaded.RunDropDatabase(username, db.Name, server.Host, server.Port, password)
				pg_reloaded.RunPsql(username, db.Name, server.Host, server.Port, db.Source.File, password)
			}
			// TODO: create schema first: RunPsql(username, db.Name, server.Host, server.Port, db.Source.File, password)
			// TODO: insert data: RunPsql(username, db.Name, server.Host, server.Port, db.Source.File, password)
		})
		if err != nil {
			fmt.Printf("Failed to start scheduler. Got error: %v \n", err)
			return err
		}
	}
	return nil
}
