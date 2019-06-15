package cmd

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var config = &Config{}
var psqlPath string
var logFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/pg_reloaded.yml)")
	rootCmd.PersistentFlags().StringVarP(&psqlPath, "psql-path", "b", "", "base project directory eg. github.com/spf13/")
	rootCmd.PersistentFlags().StringVarP(&logFile, "log-file", "l", "", "base project directory eg. github.com/spf13/")
	// TODO: rootCmd.PersistentFlags().StringVarP(&workingDir, "working-dir", "w", "", "base project directory eg. github.com/spf13/")
	rootCmd.PersistentFlags().Bool("vvv", true, "Verbose output")
	viper.BindPFlag("psql_path", rootCmd.PersistentFlags().Lookup("psql-path"))
	viper.BindPFlag("log_file", rootCmd.PersistentFlags().Lookup("log-file"))
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.SetConfigName("pg_reloaded")
		// Search config in home directory with name "pg_reloaded" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/pg_reloaded")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(config); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "pg_reloaded",
	Short: "PG Reloaded is a tool for restoring postgresql databases periodically",
	Long: `PG Reloaded is a tool for restoring postgresql databases periodically
		    for use for development and demo databases.
			More info: https://github.com/zikani03/pg_reloaded`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := Validate(config); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
