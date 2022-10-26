/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
*/

// @package: main
package main

// @import
import (
	"NFTir/server/routers"
	"NFTir/server/utils"
	"os"

	"github.com/gin-gonic/gin"
)

// @notice global variables
var (
	server      	*gin.Engine
)

// @dev Runs before main() 
func init()  {
	if (os.Getenv("GIN_MODE") != "release") {
		utils.LoadEnvVars()
	}
	
	// set up gin engine
	server = gin.Default()

	// Gin trust all proxies by default and it's not safe. Set trusted proxy to home router to to mitigate 
	server.SetTrustedProxies([]string{os.Getenv("HOME_ROUTER")})
}

// @dev Root function
func main() {
	basePath := server.Group("/v1/nnguyen6")
	routers.RouterHandler(basePath);

	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("SOURCE_IP"))
	} else {
		server.Run(":"+os.Getenv("PRODUCTION_PORT"))
	}
}