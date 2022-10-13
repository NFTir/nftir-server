/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: dynamodb.service.go provides function(s) to help client interact with AWS services
*/

/* @package services */
package utils

// import packages
import (
	"NFTir/agent/models"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

/*
@func: EstablishAwsDynamodbSession() - Establishing a connection to AWS DynamoDB

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
@func: CreateNFTirTable() - create an NFTir table on AWS DynamoDB

@params:
	- tableName string: the name of the table
	- db *dynamodb.DynamoDB: dynamodb connection
*/
func CreateNFTirTable(tableName string, db *dynamodb.DynamoDB) {
	// create table input
	tableInput := &dynamodb.CreateTableInput{
		// Table's name
		TableName: aws.String(tableName),

		// Describe the key schema for the table and indexes in AttributeDefinitions
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Name"),
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
				AttributeName: aws.String("Name"),

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

	log.Printf("Table nnguyen6_NFTir successfully created! Wating for AWS to initialize table...")
}

/*
@func: PutCollectionInput() - Putting collection items to dynamoDB table

@params:
	- tableName string: the name of the table
	- collection models.Collection: the collection item got back from NFTGo server
	- db *dynamodb.DynamoDB: dynamodb connection
*/
func PutCollectionInput(tableName string, collection models.Collection, db *dynamodb.DynamoDB) {
	log.Println(collection)
	// Get collectionAV from dynamodbattribute.marshalMap()
	collectionAttributeValue, err := dynamodbattribute.MarshalMap(collection)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	// Creating PutItemInput parameter
	collectionTableInput := &dynamodb.PutItemInput{
		Item: collectionAttributeValue,
		TableName: aws.String(tableName),
	}

	// Putting the item to the dynamoDB table
	if _, err := db.PutItem(collectionTableInput); err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
}

/*
@func: DeleteTable() - Deleting table if table already exists

@params: 
	- tableName string: the name of the table
	- db *dynamodb.DynamoDB: dynamodb connection
*/
func DeleteTable(tableName string, db *dynamodb.DynamoDB) {
	// Creating ListTableInput parameter
	listTableInput := &dynamodb.ListTablesInput{}
	
	// get the list of tableNames using ListTables API from dynamodb go sdk
	listTableOutput, err := db.ListTables(listTableInput)
	if err != nil {
		log.Fatal("Listing table error: ", err.Error())
	}

	// Checking if tableName already exists
	// Loop through TableNames array
	for _, listOutputtableName := range listTableOutput.TableNames {

		// Delete table if true
		if *listOutputtableName == tableName {
			log.Println("Table "+tableName+" already exists.")
			log.Println("Deleting table...")
			deleteTableInput := &dynamodb.DeleteTableInput{
				TableName: aws.String(tableName),
			}
		
			if _, err := db.DeleteTable(deleteTableInput); err != nil {
				log.Fatal("Deleting table error: ", err.Error())
			}
		}
	}
}