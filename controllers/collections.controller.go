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
	"github.com/jamespearly/loggly"
)

var (
	logglyHttpMessage models.HttpLogglyMessage
	logglyClient = loggly.New("NFTir")
)

/**
@func: HandleException() - loads environment varables

@param context *gin.Context - context from gin
*/
func GetStatus(context *gin.Context) {
	// Handle request methods that are not GET method
	if cMethod := context.Request.Method; cMethod != "GET" {
		// set up LogglyHttpMessage
		logglyHttpMessage = models.HttpLogglyMessage{
			Status_Code: http.StatusMethodNotAllowed,
			Method_Type: context.Request.Method,
			Source_Ip: context.ClientIP(),
			Req_Path: context.FullPath(),
		}

		// Handle Loggly
		utils.HandleLoggly(logglyClient, logglyHttpMessage, "error")

		// Abort with status
		context.AbortWithStatus(http.StatusMethodNotAllowed)
		return
	} 
	

	// Handle request with wrong path 
	if cPath := context.FullPath(); cPath != "v1/nnguyen6/status" {
		// set up LogglyHttpMessage
		logglyHttpMessage = models.HttpLogglyMessage{
			Status_Code: http.StatusNotFound,
			Method_Type: context.Request.Method,
			Source_Ip: context.ClientIP(),
			Req_Path: context.FullPath(),
		}

		// Handle Loggly
		utils.HandleLoggly(logglyClient, logglyHttpMessage, "error")
	}


	// set up LogglyHttpMessage
	logglyHttpMessage := models.HttpLogglyMessage{
		Status_Code: http.StatusOK,
		Method_Type: context.Request.Method,
		Source_Ip: context.ClientIP(),
		Req_Path: context.FullPath(),
	}

	// Handle Loggly
	utils.HandleLoggly(logglyClient, logglyHttpMessage, "info")

	// HTTP Response
	context.JSON(http.StatusOK, gin.H{
		"System-Time": time.Now(),
		"Status": http.StatusOK,
		"ClientIP": context.ClientIP(),
		"FullPath": context.FullPath(),
	  })
}

