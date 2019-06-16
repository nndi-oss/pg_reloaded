package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zikani03/pg_reloaded/pg_reloaded"
	"os"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks and validates the configuration file",
	Long:  `Checks and validates the configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := pg_reloaded.Validate(*config); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
