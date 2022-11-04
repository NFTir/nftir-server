/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: env.utils.go provides function that helps load environemnt variables
*/

// @package
package db

// @import
import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// @dev Establishing a connection to AWS DynamoDB
//
// @return *dynomadb.DynamoDB
func EstablishAwsDynamodbSession() (*dynamodb.DynamoDB) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return dynamodb.New(sess)
}