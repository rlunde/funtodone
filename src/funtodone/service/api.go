package service

import (
	"errors"
	"fmt"
	"net/http"

	"funtodone/model"
	"funtodone/session"

	"github.com/badoux/checkmail"
	"gopkg.in/gin-gonic/gin.v1"
)

/*RunService runs the main service endpoints
 * TODO:
 *   1) REST service to handle task CRUD
 *   2) User auth and session management
 *   3) REST service to handle user task collections
 *
 * TODO: convert from Gin to either net/http or (maybe) to https://github.com/go-chi/chi
 *   I don't think Gin is providing anything I want. I have never seen a complete example
 *   using middleware, which is the recommended approach for authentication. If I'm going
 *   to have to figure it out from scratch, I may as well stick to the standard library.
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
	/* session related operations: login creates a session, logout destroys one */
	r.POST("/login", LoginWithAccount)
	r.POST("/logout", Logout)

	/* all other operations require a valid session, and validation happens as a first step */
	/* I need to compose index.html dynamically, so the pieces that relate to the user (and
	   the logout link) are only shown if there is a valid session. Otherwise index.html is
		 a regular landing page with basic information and login/register links */
	/* TODO: figure out how to map all other GET urls to a standard handler that does
	   the session validation, etc. */
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
	email, password, err := getRegistrationData(c)
	if err != nil {
		c.AbortWithError(400, err)
	} else {
		fmt.Printf("RegisterAccount called with email %s, password %s\n", email, password)
	}
	//TODO: validate that account doesn't already exist
	//TODO: try to create login and save it in database
	//TODO: create a session cookie
	//TODO: return success or error message
	//TODO: on success, send email and display a verify email form
	//TODO: on error, display error message and redirect to register form
}

//LoginWithAccount -- create a new session, or return an error
func LoginWithAccount(c *gin.Context) {
	//TODO: all errors returned to JavaScript should be vague (don't say what's invalid or missing)
	email, password, err := getLoginData(c)
	if err != nil {
		c.AbortWithError(400, err)
	} else {
		fmt.Printf("LoginWithAccount called with email %s, password %s\n", email, password)
	}
	pwhash, err := model.Crypt([]byte(password))
	if err != nil {
		c.AbortWithError(400, err)
	}
	fmt.Printf("Password hash is %s\n", string(pwhash))
	//Get the user from the database if it's there
	mgr := session.GetMgr()
	user, err := model.FindUserByEmail(mgr.DbConn(), email)
	if err != nil {
		c.AbortWithError(400, err)
	}
	//verify password is correct
	if string(user.PasswordHash) != string(pwhash) {
		c.AbortWithError(400, err)
	}
	w := c.Writer
	r := c.Request
	sess, err := session.GetMgr().SessionStart(w, r)
	if err != nil {
		c.AbortWithError(400, err)
	}
	sess.Set("email", email) // puts it in the map (not yet in mongodb)
	//update the session in mongodb
	session.GetMgr().SessionUpdate(&sess)
	http.Redirect(w, r, "/", 302)
}

//Logout -- destroy a session
func Logout(c *gin.Context) {
	w := c.Writer
	r := c.Request
	session.GetMgr().SessionEnd(session.GetMgr(), w, r)
	http.Redirect(w, r, "/", 302)
	//TODO: return success or error message
	//TODO: on error, display error message and redirect back to login form
}

/*Registration - need to use BindJSON to retrieve from gin, since now posting from React as JSON struct */
type Registration struct {
	Password     string `form:"password" json:"password" binding:"required"`
	ConfPassword string `form:"confpassword" json:"confpassword" binding:"required"`
	Email        string `form:"email" json:"email" binding:"required"`
	Remember     bool   `form:"remember" json:"remember" `
}

func getRegistrationData(c *gin.Context) (email, password string, err error) {

	var json Registration
	err = c.BindJSON(&json)
	if err == nil {
		fmt.Printf("Got email: %s\n", json.Email)
	}
	err = checkmail.ValidateFormat(json.Email)
	if err != nil {
		return
	}
	err = checkmail.ValidateHost(json.Email)
	if err != nil {
		return
	}
	if json.Password != json.ConfPassword {
		err = errors.New("Password and confirm-password do not match")
	}
	email = json.Email
	password = json.Password
	return
}

/*Login - need to use BindJSON to retrieve from gin, since now posting from React as JSON struct */
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Remember bool   `form:"remember" json:"remember" `
}

func getLoginData(c *gin.Context) (username, password string, err error) {

	var json Login
	err = c.BindJSON(&json)
	if err == nil {
		fmt.Printf("Got username: %s\n", json.Username)
	}
	username = json.Username
	password = json.Password
	return
}
