/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: error.utils.go provides function that helps handle exceptions
*/

// @package
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
func HandleLoggly(logglyClient *loggly.ClientType, httpLogglyMessage models.HttpLogglyMessage, level string) error {

	// stringify struct to prepare for jearly/loggly.Send()
	stringifiedLogglyMessage, err := json.Marshal(httpLogglyMessage)
	if err != nil {return err}

	// Send message to Loggly
	err = logglyClient.Send(level, string(stringifiedLogglyMessage)); 
	if err != nil {return err}

	return nil
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
func HandleHTTPException(context *gin.Context, logglyClient *loggly.ClientType, httpFullPath string, httpMethod string) (logglyMessage *models.HttpLogglyMessage, err string) {
		if cPath := context.FullPath(); cPath != httpFullPath {
			// set up failed  LogglyHttpMessage
			logglyHttpMessage := models.HttpLogglyMessage{
				Status_Code: http.StatusNotFound,
				Method_Type: context.Request.Method,
				Source_Ip: context.ClientIP(),
				Req_Path: context.FullPath(),
			}

			return &logglyHttpMessage, "PATH";
			
		} else if cMethod := context.Request.Method; cMethod != httpMethod {
			logglyHttpMessage := models.HttpLogglyMessage{
				Status_Code: http.StatusMethodNotAllowed,
				Method_Type: context.Request.Method,
				Source_Ip: context.ClientIP(),
				Req_Path: context.FullPath(),
			}

			return &logglyHttpMessage, "METHOD";
		} 
		return nil, "";
}
