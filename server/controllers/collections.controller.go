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
	"NFTir/server/models"
	"NFTir/server/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
@func: HandleException() - loads environment varables

@param context *gin.Context - context from gin
*/
func GetStatus(context *gin.Context) {

	// load time zone
    loc, e := time.LoadLocation("EST")
	utils.HandleException(e);

	// set up LogglyHttpMessage
	logglyHttpMessage := models.HttpLogglyMessage{
		Status_Code: http.StatusOK,
		Method_Type: context.Request.Method,
		Source_Ip: context.ClientIP(),
		Req_Path: context.FullPath(),
	}

	// Handle Loggly
	utils.HandleLoggly(logglyHttpMessage, "info")

	// HTTP Response
	context.JSON(http.StatusOK, gin.H{
		"System-Time": time.Now().In(loc),
		"Status": http.StatusOK,
	  })
}

