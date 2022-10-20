/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
*/

// @package: main
package main

// Import packages
import (
	"NFTir/server/routers"
	"NFTir/server/utils"
	"os"

	"github.com/gin-gonic/gin"
)

// initialize global variables
var (
	server      	*gin.Engine
)

/* @func: init() - run before main() */
func init()  {
	if (os.Getenv("GIN_MODE") != "release") {
		utils.LoadEnvVars()
	}
	
	// set up gin engine
	server = gin.Default()

	// Gin trust all proxies by default. 192.168.1.2 typically assigned to home routers.
	server.SetTrustedProxies([]string{"192.168.1.2"})
}

/* @func main() - root function */
func main() {
	basePath := server.Group("/v1/nnguyen6")
	routers.SetupRouter(basePath);
	server.Run(os.Getenv("PORT"))
}