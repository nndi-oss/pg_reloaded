package main

import (
	"fmt"
	"github.com/nndi-oss/pg_reloaded/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
