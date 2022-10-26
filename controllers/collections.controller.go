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
	logglyClient = loggly.New("NFTir")
)

/**
@func: GetStatus() - servers the GET/status path in routers.RouterHandler

@param context *gin.Context - context from gin
*/
func GetStatus(context *gin.Context) {
	// Handle request with wrong path 
	if err := utils.HandleHTTPException(context, logglyClient); err == "PATH" {
		context.AbortWithStatus(http.StatusNotFound)
		return;
	}
	
	// Handle request methods that are not GET method
	if err := utils.HandleHTTPException(context, logglyClient); err == "METHOD" {
		context.AbortWithStatus(http.StatusMethodNotAllowed)
		return;
	}



	// set up successful LogglyHttpMessage
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

