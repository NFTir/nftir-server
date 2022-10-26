/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: gin.router.go provides function(s) related to gin's default router
*/

// @package services
package routers

// @import
import (
	"NFTir/server/controllers"

	"github.com/gin-gonic/gin"
)

// @dev handles RESTful endpoints
//
// @param routerGroup *gin.RouterGroup
func RouterHandler(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/status", controllers.GetStatus)
}