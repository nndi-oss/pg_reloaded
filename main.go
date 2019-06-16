package main

import (
	"fmt"
	"github.com/zikani03/pg_reloaded/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	  }
}
