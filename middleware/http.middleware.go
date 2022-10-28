/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: collections.service.go provides function(s) to help client process data
*/

// @package
package middleware

// @import
import (
	"regexp"

	"github.com/gin-gonic/gin"
)

// @notice search endpoint middleware
//
// @dev sanitizes passed in param, only accepts float data
//
// @return gin.HandlerFunc
func SearchMiddleware() gin.HandlerFunc {
	return func (gc *gin.Context) {
		volumeParam := gc.Param("volume_usd")

		re := regexp.MustCompile("[0-9]+[/.]?[0-9]+")
		arr := re.FindAllString(volumeParam, -1)
		if (arr == nil) {
			gc.AbortWithStatusJSON(400, gin.H{
				"error": "👀 What are you trying to do there buddy? 👀",
			})
		}
		gc.Set("sanitizedParam", arr[0])
	}
}