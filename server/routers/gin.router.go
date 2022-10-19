/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: gin.router.go provides function(s) related to gin's default router
*/

package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// init gin router
	router := gin.Default()

	// Gin trust all proxies by default. 192.168.1.2 typically assigned to home routers.
	router.SetTrustedProxies([]string{"192.168.1.2"})

	// Get request
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "Client IP": c.ClientIP(),
		})
	  })
	return router
}