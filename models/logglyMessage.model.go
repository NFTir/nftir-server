/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: logglyMessage.model.go provides LogglyMessage struct to send to loggly
*/

/** @package models */
package models

/** @struct LogglyHttpMessage - HTTP message struct to send to loggly*/
type HttpLogglyMessage struct {
	Status_Code 	uint 	`json:"status_code"`
	Method_Type 	string	`json:"method_type"`
	Source_Ip 		string 	`json:"source_ip"`
	Req_Path 		string 	`json:"req_path"`
}