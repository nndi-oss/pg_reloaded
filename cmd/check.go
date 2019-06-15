package cmd

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks and validates the configuration file",
	Long:  `Checks and validates the configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := Validate(config); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
