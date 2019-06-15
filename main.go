package main

import (
	"github.com/zikani03/pg_reloaded/cmd"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	  }
}
