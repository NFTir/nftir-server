/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: logglyMessage.model.go provides LogglyMessage struct to send to loggly
*/

// @package
package models

// @notice information realted to message struct to send to loggly
type HttpLogglyMessage struct {
	Status_Code uint 	`json:"status_code"`
	Method_Type string	`json:"method_type"`
	Source_Ip string 	`json:"source_ip"`
	Req_Path string 	`json:"req_path"`
}

// @notice information related to message struct serving GET/status
type HttpStatusMessage struct {
	Table string			`json:"table"`
	Record_Count *int64		`json:"record_count"`
}