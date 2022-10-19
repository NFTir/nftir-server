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
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
@func: HandleException() - loads environment varables

@param context *gin.Context - context from gin
*/
func GetStatus(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Client IP": context.ClientIP(),
	  })
}

