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
	"fmt"
	"log"
	"time"

	"github.com/jamespearly/loggly"
)

/*
@func PeriodicallyFetchData - fetch data in 6 hours

@params
	- logglyClient *loggly.ClientType := jearly/loggly
*/
func PeriodicallyFetchData(logglyClient *loggly.ClientType) {
	log.Println("Fetching...")
	for {
		go func () {
			// Fetching data
			NFTGoData, responseLen, err := utils.RetrieveCollectionRanking();
						
			// Processing data
			clientProcessor(logglyClient, NFTGoData, responseLen, err);
		}()
		timer1 := time.NewTimer(6*time.Hour)
		<-timer1.C
		log.Println("Refetching in 6 hours...")
	}
}

/*
@func: clientProcessor - processes parameters and sends message to Loggly 

@params
	- logglyClient *loggly.ClientType := jearly/loggly
	- NFTGoData *models.NFTGoData := models.NFTGoData
	- responseLen int: the length of the response
	- err error: error
*/
func clientProcessor(logglyClient *loggly.ClientType, NFTGoData *models.NFTGoData, responseLen int, err error)  {
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
	fmt.Printf("%+v \n", NFTGoData)
}