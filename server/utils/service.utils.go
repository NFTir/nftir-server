/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: error.utils.go provides function that helps handle exceptions
*/

/* @package initializers */
package utils

// import packages
import (
	"NFTir/server/models"
	"encoding/json"
	"log"

	"github.com/jamespearly/loggly"
)

/*
@func: HandleException() - loads environment varables

@param e error - the passed in error
*/
func HandleException(e error) {
	if (e != nil) {
		log.Fatal(e);
	}
}

func HandleLoggly(httpLogglyMessage models.HttpLogglyMessage, level string) {
	logglyClient := loggly.New("NFTir")

	// stringify struct to prepare for jearly/loggly.Send()
	stringifiedLogglyMessage, marshalErr := json.Marshal(httpLogglyMessage)
	HandleException(marshalErr)

	// Send message to Loggly
	logglyErr := logglyClient.Send(level, string(stringifiedLogglyMessage)); 
	HandleException(logglyErr)
}