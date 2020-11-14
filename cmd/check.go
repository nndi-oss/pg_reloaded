package cmd

import (
	"fmt"
	"os"

	"github.com/nndi-oss/pg_reloaded/pg_reloaded"
	"github.com/spf13/cobra"
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
