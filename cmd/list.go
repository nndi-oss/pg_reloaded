package cmd

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the configured servers and databases",
	Long:  `Lists the configured servers and databases`,
	Run: func(cmd *cobra.Command, args []string) {
		
		fmt.Println("Servers:\n===========================")
		for _, s := range config.Servers {
			fmt.Printf("Name: %s,\nHost: %s,\nPort: %s", s.Name, s.Host, s.Port)
		}
		
		fmt.Println("Databases:\n===========================")
		for _, d := range config.Databases {
			fmt.Printf("Name: %s,\nSchedule: %s,\nServer: %s,\nSource: %s",
				d.Name, d.Schedule, d.Server, d.Source.File)
		}
		fmt.Println("\n\ndone.")
	},
}