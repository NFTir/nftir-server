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
	"NFTir/agent/utils"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// initialize global variables
var (
	db *dynamodb.DynamoDB
)

/* @func: init() - run before main() */
func init()  {
	if (os.Getenv("APP_MODE") != "release") {
		utils.LoadEnvVars()
	}
	db = utils.EstablishAwsDynamodbSession()
}

/* @function main() - root function */
func main()  {
	controllers.PeriodicallyFetchData(os.Getenv("TABLE_NAME"), db)
}