/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: collections.service.go provides function(s) to help client process data
*/

// @package
package routers

// @import
import (
	"NFTir/server/controllers"
	"NFTir/server/middleware"

	"github.com/gin-gonic/gin"
)

// @notice Root struct for other methods in nft router
type NftRouter struct {
	NftController *controllers.NftirController
}


// @dev Constructor
func NftRouterConstructor(nftController *controllers.NftirController) *NftRouter {
	return &NftRouter{
		NftController: nftController,
	}
}


// @notice A medthod of nft router
// 
// @dev Declares a list of endpoints
func (nr *NftRouter) NftRoutes (rg gin.RouterGroup) {
	rg.GET("/all", nr.NftController.GetAll)
	rg.GET("/status", nr.NftController.GetStatus)
	rg.GET("/search/:volume_usd", middleware.SearchMiddleware(), nr.NftController.Search)
}