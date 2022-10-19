/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
*/

// @package: main
package main

// Import packages
import (
	"NFTir/server/routers"
	"NFTir/server/utils"
	"os"

	"github.com/jamespearly/loggly"
)

// initialize global variables
var (
	logglyClient *loggly.ClientType // loggyClient := jearly/loggly
)

/* @func: init() - run before main() */
func init()  {
	if (os.Getenv("GIN_MODE") != "release") {
		utils.LoadEnvVars()
	}
	logglyClient = loggly.New("NFTir")
}

/* @function main() - root function */
func main() {
	router := routers.SetupRouter();
	router.Run("localhost"+os.Getenv("PORT"))
}