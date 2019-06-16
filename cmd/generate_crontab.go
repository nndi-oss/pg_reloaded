package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zikani03/pg_reloaded/pg_reloaded"
	"github.com/zikani03/pg_reloaded/cron"
	"os"
)

func init() {
	rootCmd.AddCommand(generateCrontabCmd)
}

var cronTemplate = "%s \t %s"
var generateCrontabCmd = &cobra.Command{
	Use:   "generate-crontab",
	Short: "Generates a CRON Tab from PG Reloadeds configuration",
	Long:  `Generates a CRON Tab from PG Reloadeds configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		generateCrontab(args...)
	},
}

// ****** /bin/psql -U %s %s
func generateCrontab(args ...string) {
	for _, d := range config.Databases {
		v, err := cron.Parse(d.Schedule)
		if err != nil {
			fmt.Printf("Failed to parse schedule for '%s' Error %v \n", d.Name, err)
			os.Exit(1)
			break
		}
		cronSchedule := v.(*cron.SpecSchedule)
		cronStr := fmt.Sprintf("%d %d %d %d %d",
			cronSchedule.Second,
			cronSchedule.Minute,
			cronSchedule.Hour,
			cronSchedule.Dom,
			cronSchedule.Month,
		)
		server := config.GetServerByName(d.Server)
		cmdStr := pg_reloaded.DropAndRestoreUsingPsql(
			config.PsqlDir,
			server.Username,
			d.Name, 
			server.Host,
			server.Port, 
			d.Source.File, 
			server.Password,
		)

		fmt.Printf("%s\t%s\n", cronStr, cmdStr)
	}
}
