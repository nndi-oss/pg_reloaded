package cmd

import (
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
		scheduler.Run()
	},
}

func createJobsFromConfiguration(sc *cron.Cron) error {
	var server ServerConfig
	for _, db = range config.Databases {
		server = config.GetServerByName(db.Server)
		username := server.Username
		password := server.Password

		err = sc.add(db.Schedule, func() {
			if db.Source.Type == "sql" {
				RunDropDatabase(username, db.Name, server.Host, server.Port, password)
				RunPsql(username, db.Name, server.Host, server.Port, db.Source.File, password)
			}
			// TODO: create schema first: RunPsql(username, db.Name, server.Host, server.Port, db.Source.File, password)
			// TODO: insert data: RunPsql(username, db.Name, server.Host, server.Port, db.Source.File, password)
		})
		if err != nil {
			return err
		}
	}
	return nil
}
