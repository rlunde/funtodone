package service

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

/*RunService runs the main service endpoints
 * TODO:
 *   1) REST service to handle task CRUD
 *   2) User auth and session management
 *   3) REST service to handle user task collections
 *
 * Assumptions:
 *   a) use nginx to handle requests for web resources on port 80 or 443, and proxy REST requests to this port
 *   b) no need to think about user groups, roles, permissions etc. but don't make them hard to add later
 */
func RunService() {
	r := gin.Default()
	r.Static("/js", "./client/js")
	r.Static("/css", "./client/css")
	r.Static("/img", "./client/img")

	r.LoadHTMLGlob("client/*.html")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/register", RegisterAccount)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	// TODO: check r.Run for error return
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

//RegisterAccount -- create a new login
func RegisterAccount(c *gin.Context) {
	username, email, password, err := getRegistrationData(c)
	if err != nil {
		c.AbortWithError(400, err)
	} else {
		fmt.Printf("RegisterAccount called with username %s, email %s, password %s\n", username, email, password)
	}
	//TODO: validate that account doesn't already exist
	//TODO: try to create login and save it in database
	//TODO: create a session cookie
	//TODO: return success or error message
	//TODO: on success, send email and display a verify email form
	//TODO: on error, display error message and redirect to register form
}
func getRegistrationData(c *gin.Context) (username, email, password string, err error) {
	err = nil
	//sanity check input parameters
	username, found := c.GetPostForm("username")
	if !found {
		err = errors.New("RegisterAccount called but username is not set")
		return
	}
	email, found = c.GetPostForm("email")
	if !found {
		err = errors.New("RegisterAccount called but email is not set")
		return
	}
	//TODO: validate email (at least look for reasonable looking address)

	password, found = c.GetPostForm("password")
	if !found {
		err = errors.New("RegisterAccount called but password is not set")
		return
	}
	cpassword, found := c.GetPostForm("confirm-password")
	if !found {
		err = errors.New("RegisterAccount called but confirm-password is not set")
		return
	}
	if password != cpassword {
		err = errors.New("Password and confirm-password do not match")
	}
	return
}
