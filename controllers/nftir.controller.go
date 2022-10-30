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
	"NFTir/server/middleware"
	"NFTir/server/models"
	"NFTir/server/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jamespearly/loggly"
)

// @notice holds information related to NftirDao interface
type NftirController struct {
	NftirDao 	db.NftirDao
	logglyClient	*loggly.ClientType
}

// @notice constructor
func NftirControllerConstructor(nftirDao db.NftirDao, logglyClient *loggly.ClientType) *NftirController {
	return &NftirController {
		NftirDao: nftirDao,
		logglyClient: logglyClient,
	}
}

// @dev Serves the GET/all path in routers.RouterHandler
// 
// @param context *gin.Context
func (nc *NftirController) GetAll(context *gin.Context) {

	// handle loggly
	httpLogglyMessage := models.HttpLogglyMessage{
		Status_Code: http.StatusOK,
		Method_Type: context.Request.Method,
		Source_Ip: context.ClientIP(),
		Req_Path: context.FullPath(),
	}
	if err := utils.HandleLoggly(nc.logglyClient, httpLogglyMessage, "info"); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"loggly_error": err.Error()})
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
	// handle loggly
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

func (nc *NftirController) Search(context *gin.Context) {
	// get volume param
	volumeParam := context.MustGet("sanitizedParam").(string);

	// convert from string to float 32
	volume64, err := strconv.ParseFloat(volumeParam, 32)
	if err != nil {context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Cannot parse float volume param"})}
	volume32 := float32(volume64)

	// Get collections
	collections, err := nc.NftirDao.GetCollectionsGreaterThan(&volume32)
	if err != nil {context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})}

	// HTTP 200 Response
	context.JSON(http.StatusOK, collections)
}


// @notice HTTP endpoints
func (nc *NftirController) FetchCollectionsRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/all", nc.GetAll)
	routerGroup.GET("/status", nc.GetStatus)
	routerGroup.GET("/search/:volume_usd", middleware.SearchMiddleware(), nc.Search)
}