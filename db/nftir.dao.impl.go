/*
	@Author: Logan (Nam) Nguyen
	@Course: SUNY Oswego - CSC 482
	@Instructor: Professor James Early
	@Purpose: env.utils.go provides function that helps load environemnt variables
*/

// @package
package db

// @import
import (
	"NFTir/server/models"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// @notice Holds information related to dynamodb
type NftirDaoImpl struct {
	dynamodb *dynamodb.DynamoDB
}

// @dev constructor
func NftirDaoConstructor(dynamodb *dynamodb.DynamoDB) NftirDao {
	return &NftirDaoImpl {
		dynamodb: dynamodb,
	}
}

// @notice Gets all the collections from dynamoDB table
// 
// @return *[]models.Collection
// 
// @return error
func (ndi *NftirDaoImpl) GetAll() (*[]models.Collection, error) {

	// Scan the dynamoDB `TABLE_NAME` table
	collectionSO, err := ndi.dynamodb.Scan(&dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	})
	if err != nil {return nil, err}

	
	// dynamic slice
	collections := make([]models.Collection, *collectionSO.Count)
	
	// process collectionAV *dynamodb.ScanOutput
	for index, collectionsAV := range collectionSO.Items {
		collection := models.Collection{}

		if err = dynamodbattribute.UnmarshalMap(collectionsAV, &collection); err != nil {
			return nil, err
		}

		collections[index] = collection
	}

	return &collections, nil
}

// @dev Gets the total number of collections stored in the NFTier dynamoDB table
// 
// @return *models.HttpStatusMessage
// 
// @return error
func (ndi *NftirDaoImpl) GetStatus() (*models.HttpStatusMessage, error) {
	return nil, nil
}

// @dev Gets a subset of collections based on the params passed in
// 
// @param volumeUsd *float64
// 
// @param marketCapUsd *float64
// 
// @return *[]models.Colelction
// 
// @return error
func (ndi *NftirDaoImpl) GetCollectionsBy(volumeUsd *float64, marketCapUsd *float64) (*[]models.Collection, error) {
	return nil, nil;
}