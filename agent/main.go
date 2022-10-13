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

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jamespearly/loggly"
)

// initialize global variables
var (
	logglyClient *loggly.ClientType // loggyClient := jearly/loggly
	tableName string
	db *dynamodb.DynamoDB
)

/* @func: init() - run before main() */
func init()  {
	initializers.LoadEnvVars();
	logglyClient = loggly.New("NFTir")
	tableName = "nnguyen6_NFTir_v1"
	db = utils.EstablishAwsDynamodbSession()
}

/* @function main() - root function */
func main()  {
	controllers.PeriodicallyFetchData(logglyClient, tableName, db);
}