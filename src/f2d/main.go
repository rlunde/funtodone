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
	service.AuBo = authboss.New() // Usually store this globally
	AuBoMountPath = "/authboss"
	service.AuBo.LogWriter = os.Stdout

	if err := service.AuBo.Init(); err != nil {
		// Handle error, don't let program continue to run
		log.Fatalln(err)
	}

	// Make sure to put authboss's router somewhere

	service.RunService()
}
