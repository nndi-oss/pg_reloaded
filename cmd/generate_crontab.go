package cmd

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
		cronStr = string(cron.ParseSchedule(d.Scheduler))
		server := config.GetServerByName(d.Server)
		cmdStr := dropAndRestoreUsingPsql(
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
