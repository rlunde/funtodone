package main

import (
	"fmt"
	"funtodone/service"
	//"log"
	//"net/http"
	//"os"
)

/*
 * This is the main routine of the fun2done REST service. While it
 * is possible to serve the HTML, JS, and CSS files from this, it's
 * probably better to serve all those via nginx and use nginx proxy
 * to map the REST URLs to the port this service is running on.
 *
 * Currently:
 *   http://localhost/f2d/ -- maps to ../client
 *   cd .. ; ./build -- build this service
 *   ./funtodone -- run this service
 */

var globalSessions *service.SessionManager

//  initialize the session manager (init is run automatically)
func init() {
	var err error
	globalSessions, err = service.NewManager("memory", "gosessionid", 3600)
	if err != nil {
		fmt.Printf("Error creating session manager: %s", err.Error())
	}
}
func main() {
	service.RunService()
}
