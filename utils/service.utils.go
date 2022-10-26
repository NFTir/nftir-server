/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: error.utils.go provides function that helps handle exceptions
*/

// @package initializers
package utils

// @import
import (
	"NFTir/server/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamespearly/loggly"
)

// @dev Handle interacting with Loggly
//
// @param logglyClient *loggly.ClientType
//
// @param httpLogglyMessage models.HttpLogglyMessage
//
// @param level string
func HandleLoggly(logglyClient *loggly.ClientType, httpLogglyMessage models.HttpLogglyMessage, level string) {

	// stringify struct to prepare for jearly/loggly.Send()
	stringifiedLogglyMessage, marshalErr := json.Marshal(httpLogglyMessage)
	HandleException(marshalErr)

	// Send message to Loggly
	logglyErr := logglyClient.Send(level, string(stringifiedLogglyMessage)); 
	HandleException(logglyErr)
}


// @dev Handdle error exception
//
// @param e error - the passed in error
func HandleException(e error) {
	if (e != nil) {
		log.Fatal(e);
	}
}


// @dev Handdle HTTP exception
// 
// @param context *gin.Context
// 
// @param logglyClient *loggly.ClientType
// 
// @return err string
func HandleHTTPException(context *gin.Context, logglyClient *loggly.ClientType) (err string) {
		if cPath := context.FullPath(); cPath != "/v1/nnguyen6/status" {
			// set up failed  LogglyHttpMessage
			logglyHttpMessage := models.HttpLogglyMessage{
				Status_Code: http.StatusNotFound,
				Method_Type: context.Request.Method,
				Source_Ip: context.ClientIP(),
				Req_Path: context.FullPath(),
			}

			// Handle Loggly
			HandleLoggly(logglyClient, logglyHttpMessage, "error")

			return "PATH";
		} else if cMethod := context.Request.Method; cMethod != "GET" {
			logglyHttpMessage := models.HttpLogglyMessage{
				Status_Code: http.StatusMethodNotAllowed,
				Method_Type: context.Request.Method,
				Source_Ip: context.ClientIP(),
				Req_Path: context.FullPath(),
			}
	
			// Handle Loggly
			HandleLoggly(logglyClient, logglyHttpMessage, "error")
	
			return "METHOD";
		} 
		return;
}

