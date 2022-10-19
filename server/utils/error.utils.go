/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: error.utils.go provides function that helps handle exceptions
*/

/* @package initializers */
package utils

// import packages
import "log"

/*
@func: HandleException() - loads environment varables

@param e error - the passed in error
*/
func HandleException(e error, s string) {
	if (e != nil) {
		log.Fatal(s);
	}
}