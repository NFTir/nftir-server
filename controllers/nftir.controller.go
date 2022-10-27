/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: collections.service.go provides function(s) to help client process data
*/

// @package
package controllers

// @import
import (
	"NFTir/server/db"
	"NFTir/server/models"
	"NFTir/server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jamespearly/loggly"
)

// @notice holds information related to NftirDao interface
type NftirController struct {
	NftirDao 		db.NftirDao
	logglyClient	*loggly.ClientType
}

// @notice constructor
func NftirControllerConstructor(nftirDao db.NftirDao, logglyClient	*loggly.ClientType) *NftirController {
	return &NftirController {
		NftirDao: nftirDao,
		logglyClient: logglyClient,
	}
}

// @dev Serves the GET/all path in routers.RouterHandler
// 
// @param context *gin.Context
func (nc *NftirController) GetAll(context *gin.Context) {

	// handle request with wrong path 
	if httpLogglyMessage, err := utils.HandleHTTPException(context, nc.logglyClient, context.FullPath(), context.Request.Method); err == "PATH" {
		if err := utils.HandleLoggly(nc.logglyClient, *httpLogglyMessage, "error"); err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"loggly_error": err.Error()})
			return;
		}
		context.AbortWithStatus(http.StatusNotFound)
		return;
	}
	
	// handle request methods that are not GET method
	if httpLogglyMessage, err := utils.HandleHTTPException(context, nc.logglyClient, context.FullPath(), context.Request.Method); err == "METHOD" {
		if err := utils.HandleLoggly(nc.logglyClient, *httpLogglyMessage, "loggly_error"); err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		context.AbortWithStatus(http.StatusMethodNotAllowed)
		return;
	}

	// handle successful loggly
	httpLogglyMessage := models.HttpLogglyMessage{
		Status_Code: http.StatusOK,
		Method_Type: context.Request.Method,
		Source_Ip: context.ClientIP(),
		Req_Path: context.FullPath(),
	}
	if err := utils.HandleLoggly(nc.logglyClient, httpLogglyMessage, "info"); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"loggly_error": err})
		return;
	}

	// retrieve collections by implementing GetAll() from dao impl
	collections, err := nc.NftirDao.GetAll()
	if err!= nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return;
	}

	// HTTP 200 Response
	context.JSON(http.StatusOK, collections)
}


// @dev Serves the GET/status path in routers.RouterHandler
// 
// @param context *gin.Context
func (nc *NftirController) GetStatus(context *gin.Context) {
	// handle request with wrong path 
	if httpLogglyMessage, err := utils.HandleHTTPException(context, nc.logglyClient, context.FullPath(), context.Request.Method); err == "PATH" {
		if err := utils.HandleLoggly(nc.logglyClient, *httpLogglyMessage, "error"); err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"loggly-error": err.Error()})
			return;
		}
		context.AbortWithStatus(http.StatusNotFound)
		return;
	}
	
	// handle request methods that are not GET method
	if httpLogglyMessage, err := utils.HandleHTTPException(context, nc.logglyClient, context.FullPath(), context.Request.Method); err == "METHOD" {
		if err := utils.HandleLoggly(nc.logglyClient, *httpLogglyMessage, "loggly-error"); err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		context.AbortWithStatus(http.StatusMethodNotAllowed)
		return;
	}

	// handle successful loggly
	httpLogglyMessage := models.HttpLogglyMessage{
		Status_Code: http.StatusOK,
		Method_Type: context.Request.Method,
		Source_Ip: context.ClientIP(),
		Req_Path: context.FullPath(),
	}
	if err := utils.HandleLoggly(nc.logglyClient, httpLogglyMessage, "info"); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"loggly_error": err})
		return;
	}

	// retrieve table status by implementing GetStatus() from dao impl
	httpStatusMessage, err := nc.NftirDao.GetStatus()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return;
	}

	// HTTP 200 Response
	context.JSON(http.StatusOK, httpStatusMessage)
}


// @notice HTTP endpoints
func (nc *NftirController) FetchCollectionsRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/all", nc.GetAll)
	routerGroup.GET("/status", nc.GetStatus)
}