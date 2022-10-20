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

	// Gin trust all proxies by default and it's not safe. Set trusted proxy to home router to to mitigate 
	server.SetTrustedProxies([]string{os.Getenv("HOME_ROUTER")})
}

/* @func main() - root function */
func main() {
	basePath := server.Group("/v1/nnguyen6")
	routers.SetupRouter(basePath);
	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("SOURCE_IP"))
	} else {
		server.Run(":"+os.Getenv("PORT"))
	}
}