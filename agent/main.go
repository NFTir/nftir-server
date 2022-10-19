/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
*/

// @package: main
package main

// Import packages
import (
	"NFTir/agent/controllers"
	"NFTir/agent/initializers"
	"NFTir/agent/utils"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jamespearly/loggly"
)

// initialize global variables
var (
	logglyClient *loggly.ClientType // loggyClient := jearly/loggly
	db *dynamodb.DynamoDB
)

/* @func: init() - run before main() */
func init()  {
	if (os.Getenv("APP_ENV") != "production") {
		initializers.LoadEnvVars()
	}
	logglyClient = loggly.New("NFTir")
	db = utils.EstablishAwsDynamodbSession()
}

/* @function main() - root function */
func main()  {
	controllers.PeriodicallyFetchData(logglyClient, os.Getenv("TABLE_NAME"), db);
}