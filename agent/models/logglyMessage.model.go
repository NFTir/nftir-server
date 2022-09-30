/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: logglyMessage.model.go provides LogglyMessage struct to send to loggly
*/

/* @package models */
package models

/* @struct LogglyMessage - message struct to send to loggly*/
type LogglyMessage struct {
	Request_Status 	string	`json:"request_status"`
	Data_Length 	int 	`json:"data_length"`
}