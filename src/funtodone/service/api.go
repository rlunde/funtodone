package service

import "gopkg.in/gin-gonic/gin.v1"

/*
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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
