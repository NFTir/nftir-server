/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
*/

// @package: main
package main

import (
	"NFTir/server/controllers"
	"NFTir/server/db"
	"NFTir/server/utils"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/jamespearly/loggly"
)

// @import

// @notice global variables
var (
	server      	*gin.Engine
	logglyClient	*loggly.ClientType
	dbconn 			*dynamodb.DynamoDB
	nd				db.NftirDao
	nc 				*controllers.NftirController
)

// @dev Runs before main() 
func init()  {
	// load env variables
	if (os.Getenv("GIN_MODE") != "release") {utils.LoadEnvVars()}
	
	// set up gin engine
	server = gin.Default()

	// Gin trust all proxies by default and it's not safe. Set trusted proxy to home router to to mitigate 
	server.SetTrustedProxies([]string{os.Getenv("HOME_ROUTER")})

	// init jearly/loggly
	logglyClient = loggly.New("NFTir")

	// init db connection
	dbconn = utils.EstablishAwsDynamodbSession()

	// init nftirDao interface
	nd = db.NftirDaoConstructor(dbconn)

	// init nftController
	nc = controllers.NftirControllerConstructor(nd, logglyClient)
}

// @dev Root function
func main() {
	// base path
	basePath := server.Group("/v1/nnguyen6")

	// init handlers
	nc.FetchCollectionsRoutes(basePath)

	// run server
	if (os.Getenv("GIN_MODE") != "release") {
		server.Run(os.Getenv("SOURCE_IP"))
	} else {
		server.Run(":"+os.Getenv("PRODUCTION_PORT"))
	}
}