package cmd

import (
	"fmt"
	"github.com/nndi-oss/pg_reloaded/cron"
	"github.com/nndi-oss/pg_reloaded/pg_reloaded"
	"github.com/spf13/cobra"
)

var scheduler = cron.New()
var shutdownCh chan (struct{})

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the PG Reloaded scheduler/cron daemon",
	Long:  `Starts the PG Reloaded scheduler/cron daemon`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := createJobsFromConfiguration(scheduler); err != nil {
			return err
		}

		scheduler.Start()
		select {
		// TODO: handle signals case s := <-signalCh:
		case <-shutdownCh:
			return nil
		}
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
				pg_reloaded.RunDropDatabase(
					config.PsqlDir,
					username,
					db.Name,
					server.Host,
					server.Port,
					password,
				)
				pg_reloaded.RunRestoreDatabase(
					config.PsqlDir,
					username,
					db.Name,
					server.Host,
					server.Port,
					db.Source.File,
					password,
				)
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
