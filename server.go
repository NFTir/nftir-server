/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
*/

// @package: main
package main

// @import
import (
	"NFTir/server/controllers"
	"NFTir/server/dao"
	"NFTir/server/db"
	"NFTir/server/routers"
	"NFTir/server/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jamespearly/loggly"
)

// @notice global variables
var (
	server		*gin.Engine
	logglyClient	*loggly.ClientType
	nr				*routers.NftRouter
)

// @dev Runs before main() 
func init()  {
	// load env variables
	if (os.Getenv("GIN_MODE") != "release") {utils.LoadEnvVars()}
	
	// set up gin engine - Default With the Logger and Recovery middleware already attached
	server = gin.Default()

	// Gin trust all proxies by default and it's not safe. Set trusted proxy to home router to to mitigate 
	server.SetTrustedProxies([]string{os.Getenv("HOME_ROUTER")})

	// init jearly/loggly
	logglyClient = loggly.New("NFTir")

	// init db connection
	dbconn := db.EstablishAwsDynamodbSession()

	// init nftirDao interface
	nd := dao.NftirDaoConstructor(dbconn)

	// init nftController
	nc := controllers.NftirControllerConstructor(nd, logglyClient)

	// init nftRouter 
	nr = routers.NftRouterConstructor(nc)
	
}

// @dev Root function
func main() {
	server.HandleMethodNotAllowed = true

	// base path
	basePath := server.Group("nnguyen6")

	// init handlers
	nr.NftRoutes(*basePath)

	// run server
	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("SOURCE_IP"))
	} else {
		server.Run(":"+os.Getenv("PRODUCTION_PORT"))
	}
}