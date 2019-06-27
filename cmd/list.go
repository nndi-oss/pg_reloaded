package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the configured servers and databases",
	Long:  `Lists the configured servers and databases`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Servers:")
		fmt.Println("===========================")
		for _, s := range config.Servers {
			fmt.Println("Name: ", s.Name)
			fmt.Println("Host: ", s.Host)
			fmt.Println("Port: ", s.Port)
			fmt.Println("Username: ", s.Username)
		}
		fmt.Println("\n")
		fmt.Println("Databases:")
		fmt.Println("===========================")
		for _, d := range config.Databases {
			fmt.Println("Name:", d.Name)
			fmt.Println("Schedule:", d.Schedule)
			fmt.Println("Server:", d.Server)
			fmt.Println("Source:", d.Source.File)
		}
		fmt.Println("\n\ndone.")
	},
}
