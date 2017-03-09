package main

import (
	"funtodone/service"
	"log"
	"os"

	authboss "gopkg.in/authboss.v1"
	//"log"
	//"net/http"
	//"os"
)

/*
 */
func main() {
	ab := authboss.New() // Usually store this globally
	ab.MountPath = "/authboss"
	ab.LogWriter = os.Stdout

	if err := ab.Init(); err != nil {
		// Handle error, don't let program continue to run
		log.Fatalln(err)
	}

	// Make sure to put authboss's router somewhere

	service.RunService()
}
