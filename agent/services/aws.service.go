/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: dynamodb.service.go provides function(s) to help client interact with AWS services
*/

/* @package services */
package services

// import packages
import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/*
@func: EstablishAwsDynamodbSession - Establishing a connection to AWS DynamoDB

@return
	- db *dynamodb.DynamoDB
*/
func EstablishAwsDynamodbSession() (*dynamodb.DynamoDB) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return dynamodb.New(sess);
}

/*
@func: CreateNFTirTable - create an NFTir table on AWS DynamoDB
NOTE: This script is run ONLY once in main.go to create the table
*/
func CreateNFTirTable() {
	// establish the db connection
	db := EstablishAwsDynamodbSession()

	// table's name
	tableName := "nnguyen6_NFTir"

	// create table input
	tableInput := &dynamodb.CreateTableInput{
		// Table's name
		TableName: aws.String(tableName),

		// Describe the key schema for the table and indexes in AttributeDefinitions
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Contracts"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Volume_usd"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement {
			// Partition key Contracts will be unique
			{
				AttributeName: aws.String("Contracts"),

				// KeyType = HASH - The attribute that DynamoDB will use to partition the data onto one of its many storage nodes
				// If ONLY partition key is specified and NOT a sort key, all records must have a unique partition key value.
				
				// KeyType = RANGE - The secondary key that can be optionally decided to use alongside the Partition Key
				
				// Traditional SQL Primary Key can be either just a Partition Key, OR Partition Key + Sort Key (a.k.a composite primary key)

				// With composite primary key, data can be stored with the SAME parition key value but a different sort key value
				// Read more on partition and sort key https://www.beabetterdev.com/2022/02/07/dynamodb-partition-key-vs-sort-key/
				KeyType: aws.String("HASH"), 
			},
			
			// Sort key Volume_usd will be used to perform "range-like" querires (sorting, equality, comparision, etc.)
			{
				AttributeName: aws.String("Volume_usd"),
				KeyType: aws.String("RANGE"),
			},
		},
		
		// Throughput settings
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(100),
		},
	}

	// Create the table
	if _, err := db.CreateTable(tableInput); err != nil {
		// if err is not nil, fatalize the process
		log.Fatal("Creating table error: ", err.Error())
	}

	log.Printf("nnguyen6_NFTir table successfully created!")
}