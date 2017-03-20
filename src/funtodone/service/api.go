package service

import (
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
	r.Static("resources", "./resources")

	r.LoadHTMLGlob("resources/views/gin-gonic/*")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		data := layoutData(c.Writer, c.Request)
		c.HTML(http.StatusOK, "index.tmpl", data)
	})

	initAuthBossRoute(r)
	// Using authboss is much more complicated than this -- see:
	// https://github.com/go-authboss/authboss/wiki/Integration-with-gin-gonic
	// http.Handle("/authboss", ab.NewRouter())
	// http.ListenAndServe(":8082", nil) // main service endpoints are on 8080
	// if authboss is on another port, we may have to enble CORS?
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
