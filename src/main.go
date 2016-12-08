package main

import (
	"fmt"
	//"log"
	"./cli/cmd"
	//"net/http"
	"os"
)

/*
For now, we'll just do a command line app. Later on, we'll
add a comman line option that will start a web service using Gin.

See https://github.com/spf13/cobra for how to generate code using
cobra/viper and how to edit it.
*/
func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
