/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: gin.router.go provides function(s) related to gin's default router
*/

/** @package services */
package routers

// import packages
import (
	"NFTir/server/controllers"

	"github.com/gin-gonic/gin"
)

/**
@func: SetupRouter() - set up RESTful endpoints
*/
func SetupRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/status", controllers.GetStatus)
}