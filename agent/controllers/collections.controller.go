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
	"NFTir/agent/models"
	"NFTir/agent/utils"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jamespearly/loggly"
)

/*
@func: PeriodicallyFetchData - fetch data in 6 hours

@params
	- logglyClient *loggly.ClientType := jearly/loggly
	- tableName string: the name of the table
	- db *dynamodb.DynamoDB: dynamodb connection
*/
func PeriodicallyFetchData(logglyClient *loggly.ClientType, tableName string, db *dynamodb.DynamoDB) {
	for {
		setUpTableAsync(tableName, db);
		fetchDataAsync(logglyClient, tableName, db)
	}
}

/*
@func: setUpTableAsync - set up dynamodb table
@params:
	- tableName string: the name of the table
	- db *dynamodb.DynamoDB: dynamodb connection
@TODO: Implement real ASYNC/AWAIT
*/
func setUpTableAsync(tableName string, db *dynamodb.DynamoDB) {
	log.Println("Starting polling process...")
	utils.DeleteTable(tableName, db)
	time1 := time.NewTimer(10*time.Second)
	<- time1.C
	log.Println("Creating new table "+tableName+"...")
	utils.CreateNFTirTable(tableName, db)
	timer2 := time.NewTimer(10*time.Second)
	<-timer2.C
	log.Println("Finished initializing table "+tableName+".")
}

/*
@func: fetchDataAsync - fetching data from NFTGo server
@params:
	- logglyClient *loggly.ClientType := jearly/loggly
	- tableName string: the name of the table
	- db *dynamodb.DynamoDB: dynamodb connection
*/
func fetchDataAsync(logglyClient *loggly.ClientType, tableName string, db *dynamodb.DynamoDB) {
	go func () {
		// Fetching data
		NFTGoData, responseLen, err := utils.RetrieveCollectionRanking();
					
		// Processing data
		clientProcessor(tableName, logglyClient, NFTGoData, responseLen, err, db);
	}()
	timer := time.NewTimer(6*time.Hour)
	<-timer.C
}

/*
@func: clientProcessor - processes parameters and sends message to Loggly 

@params
	- logglyClient *loggly.ClientType := jearly/loggly
	- NFTGoData *models.NFTGoData := models.NFTGoData
	- responseLen int: the length of the response
	- err error: error
*/
func clientProcessor(tableName string, logglyClient *loggly.ClientType, NFTGoData *models.NFTGoData, responseLen int, err error, db *dynamodb.DynamoDB)  {
	
	// if err := controllers.RetrieveCollectionRanking() != nil, send a failed message to loggly then fatalize the process with err message and a call to os.Exit(1)
	if err != nil {
		// init LogglyMessage struct with failed Request_Status and Data_Length equals to 0
		LogglyMessage := models.LogglyMessage{Request_Status: "failure", Data_Length: 0}
		// stringify struct to prepare for jearly/loggly.Send()
		stringifiedLogglyMessage, marshalErr := json.Marshal(LogglyMessage)
		if marshalErr != nil {
			log.Fatal(marshalErr)
		}

		// if logglyError := jearly/loggly.Send() != nil, fatalize the process with the loglly a call to os.Exit(1)
		if logglyErr := logglyClient.Send("error", string(stringifiedLogglyMessage)); logglyErr != nil {
			log.Fatal(logglyErr)
		}

		// fatilze the error
		log.Fatal(err)
	}

	// if err == nil, prepare a LogglyMessage struct with successful Request_Staus and Data_Length 
	LogglyMessage := models.LogglyMessage{Request_Status: "success", Data_Length: responseLen}
	
	// stringify struct to prepare for jearly/loggly.Send()
	stringifiedLogglyMessage, marshalErr := json.Marshal(LogglyMessage)
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}
	
	// if logglyError := jearly/loggly.Send() != nil, fatalize the process with the loglly a call to os.Exit(1)
	if logglyErr := logglyClient.Send("info", string(stringifiedLogglyMessage)); logglyErr != nil {
		log.Fatal((logglyErr))
	}
	
	// Print NFTGoData := controllers.RetrieveCollectionRanking() to the console
	// fmt.Printf("%+v \n", NFTGoData)

	// Pushing NFTGo collections to table
	log.Println("Pushing NFTGo collections to table "+tableName)
	for _, collection := range NFTGoData.Collections {
		utils.PutCollectionInput(tableName, collection, db)
	}
	log.Println("Successfully push collection items to table "+tableName)
	log.Println("Refetching in 6 hours...")
}

