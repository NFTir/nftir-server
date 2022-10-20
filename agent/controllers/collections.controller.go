/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: collections.service.go provides function(s) to help client process data
*/

/* @package services */
package controllers

// import packages
import (
	"NFTir/agent/utils"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/*
@func: PeriodicallyFetchData() - fetch data in 6 hours

@params
	- logglyClient *loggly.ClientType := jearly/loggly
	- tableName string: the name of the table
	- db *dynamodb.DynamoDB: dynamodb connection
*/
func PeriodicallyFetchData(tableName string, db *dynamodb.DynamoDB) {
	for { // infinite loop. Gap time = 6 hours
		utils.SetUpTableAsync(tableName, db);
		utils.FetchDataAsync(tableName, db)
	}
}
