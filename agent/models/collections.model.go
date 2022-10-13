/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: collection.model.go provides NFTGoResponse struct which mimics the data received from the API request
*/

/* @package models */
package models

/* @struct Collection - Each item in collections slice in @struct/NFTGoResponse */
type Collection struct {
	Last_updated 							uint 		`json:"last_updated"`
	Blockchain 								string 		`json:"blockchain"`
	Name 									string 		
	Slug 									string 		`json:"slug"`
	Opensea_slug 							string 		`json:"opensea_slug"`
	Description 							string  	`json:"description"`
	Official_website_url 					string 		`json:"official_website_url"`
	Opensea_url 							string 		`json:"opensea_url"`
	Logo 									string 		`json:"logo"`
	Contracts 								[]string 	`json:"contracts"`
	Contract_type 							string 		`json:"contract_type"`
	Categories 								[]string 	`json:"categories"`
	Discord_url 							string 		`json:"discord_url"`
	Instagram_url 							string 		`json:"instagram_url"`
	Twitter_url 							string 		`json:"twitter_url"`
	Telegram_url 							string 		`json:"telegram_url"`
	Has_rarity 								bool 		`json:"has_rarity"`
	Market_cap_usd 							float32 	`json:"market_cap_usd"`
	Market_cap_change_percentage 			float64 	`json:"market_cap_change_percentage"`
	Volume_usd 								float32 	
	Floor_price_usd 						float32 	`json:"floor_price_usd"`
	Whale_num 								uint 		`json:"whale_num"`
	Holder_num 								uint 		`json:"holder_num"`
}

/* @struct NFTGoResponse - final HTTP response from NFTGo API request*/
type NFTGoData struct {
	Total			uint			`json:"total"`
	Collections		[]Collection	`json:"collections"`
}