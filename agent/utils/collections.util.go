/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: collections.controller.go provides function to make an API request to NFTGo server
*/

/* @package models */
package utils

import (
	"NFTir/agent/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/*
@func RtrieveCollectionRanking() - Controller that makes an http request to NFTGoAPI server

@return
	- NFTGoData *models.NFTGoData := models.NFTGoData
	- responseLen int: the length of the response
	- err error: error
*/
func retrieveCollectionRanking() (NFTGoData *models.NFTGoData, responseLen int, err error) {
	log.Println("Fetching data from NFTGo server...")
	// retrieve env vars
	NFTGO_API_URL := os.Getenv("NFTGO_API_URL")
	NFTGO_API_KEY := os.Getenv("NFTGO_API_KEY")

	// Init new http request
	req, err := http.NewRequest("GET", NFTGO_API_URL, nil)
	if err != nil {
		return NFTGoData, 0, err
	}

	// Add neccessary key/value pairs to request header
	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", NFTGO_API_KEY)

	// Make the request to NFTFo server to fetch data
	res, err := http.DefaultClient.Do(req)
	if (err != nil) {
		return NFTGoData, 0, err
	}

	// make sure to close after everything is done
	defer res.Body.Close()

	// process res.body using fuction ReadAll from ioutil package
	body, _ := ioutil.ReadAll(res.Body)

	// parse json from []byte to JSON
	json.Unmarshal(body, &NFTGoData)
	responseLen = len(body)

	// return NFTGoData and nil for err
	return NFTGoData, responseLen, nil
}


/*
@func: clientProcessor() - processes parameters and sends message to Loggly 

@params
	- logglyClient *loggly.ClientType := jearly/loggly
	- NFTGoData *models.NFTGoData := models.NFTGoData
	- responseLen int: the length of the response
	- err error: error
*/
func clientProcessor(tableName string, NFTGoData *models.NFTGoData, responseLen int, err error, db *dynamodb.DynamoDB)  {
	
	// if err := controllers.RetrieveCollectionRanking() != nil, send a failed message to loggly then fatalize the process with err message and a call to os.Exit(1)
	if err != nil {
		// init LogglyMessage struct with failed Request_Status and Data_Length equals to 0
		LogglyMessage := models.LogglyMessage{Request_Status: "failure", Data_Length: 0}

		// Handle Loggly
		HandleLoggly(LogglyMessage, "error")
		
		// fatilze the error
		log.Fatal(err)
	}

	// if err == nil, prepare a LogglyMessage struct with successful Request_Staus and Data_Length 
	LogglyMessage := models.LogglyMessage{Request_Status: "success", Data_Length: responseLen}

	// Handle Loggly
	HandleLoggly(LogglyMessage, "info")
	
	// Pushing NFTGo collections to table
	log.Println("Pushing NFTGo collections to table "+tableName)
	for _, collection := range NFTGoData.Collections {
		PutCollectionInput(tableName, collection, db)
	}
	log.Println("Successfully push collection items to table "+tableName)
	
	log.Println("Refetching in 6 hours...")
}


/*
@func: fetchDataAsync() - fetching data from NFTGo server
@params:
	- logglyClient *loggly.ClientType := jearly/loggly
	- tableName string: the name of the table
	- db *dynamodb.DynamoDB: dynamodb connection
*/
func FetchDataAsync(tableName string, db *dynamodb.DynamoDB) {
	go func () {
		// Fetching data from NFTGo server
		NFTGoData, responseLen, err := retrieveCollectionRanking();
					
		// Processing data and push collection items to dynamodb table tableName
		clientProcessor(tableName, NFTGoData, responseLen, err, db);
	}()
	timer := time.NewTimer(6*time.Hour)
	<-timer.C
}
