/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: collections.controller.go provides function to make an API request to NFTGo server
*/

/* @package models */
package controllers

import (
	"NFTir/agent/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

/*
@function RtrieveCollectionRanking - Controller that makes an http request to NFTGoAPI server

@return
	- NFTGoData *models.NFTGoData := models.NFTGoData
	- responseLen int: the length of the response
	- err error: error
*/
func RetrieveCollectionRanking() (NFTGoData *models.NFTGoData, responseLen int, err error) {
	// retrieve env vars
	NFTGO_API_URL := os.Getenv("NFTGO_API_URL")
	NFTGO_API_KEY := os.Getenv("NFTGO_API_KEY")

	// Init new http request
	req, err := http.NewRequest("GET", NFTGO_API_URL, nil)
	if err != nil {
		return NFTGoData, 0, err
	}

	// Add neccessary key/value pairs to request header
	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", NFTGO_API_KEY)

	// Make the request to NFTFo server to fetch data
	res, err := http.DefaultClient.Do(req)
	if (err != nil) {
		return NFTGoData, 0, err
	}

	// make sure to close after everything is done
	defer res.Body.Close()

	// process res.body using fuction ReadAll from ioutil package
	body, _ := ioutil.ReadAll(res.Body)

	// parse json from []byte to JSON
	json.Unmarshal(body, &NFTGoData)
	responseLen = len(body)

	// return NFTGoData and nil for err
	return NFTGoData, responseLen, nil
}