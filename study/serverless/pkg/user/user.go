package user

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/gusw/go_main/study/serverless/pkg/validators"
)

var(
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorFailedToFetchRecord = "failed to fetch record"
	ErrorInvalidUserData = "invalid user data"
	ErrorInvalidEmail = "invalid email"
	ErrorCouldNotMarshalItem = "could not marshal item"
	ErrorCouldNotDeleteItem = "could not delete item"
	ErrorCouldNotDynamoPutItem = "could not dynamo put item"
	ErrorUserAlreadyExists = "user.User already exists"
	ErrorUserDoesNotExist = "user.User does not exist"
)

type User struct{
	Email 		string	`json:"email"`
	FirstName	string 	`json:"firstName"`
	LastName	string 	`json:"lastName"`
}

func _handleError(errorType string) (*User, error){
	return nil, errors.New(errorType)
}

func _handleErrors(errorType string) (*[]User, error){
	return nil, errors.New(errorType)
}

func _generateUserKey(email string)(map[string]*dynamodb.AttributeValue){
	return map[string]*dynamodb.AttributeValue{
			"email":{
				S: aws.String(email),
			},
		}
}

func _checkInvalidUserData(req events.APIGatewayProxyRequest, user *User)(bool){
	if err := json.Unmarshal([]byte(req.Body), &user); err != nil {
		return true
	}
	return false
}

func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*User, error){

	input := &dynamodb.GetItemInput{
		Key: _generateUserKey(email),
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.GetItem(input)
	if err != nil {
		return _handleError(ErrorFailedToFetchRecord)
	}

	item := new(User)

	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return _handleError(ErrorFailedToUnmarshalRecord)
	}

	return item, nil
}

func FetchUsers(tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*[]User, error){
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.Scan(input)
	if err != nil {
		return _handleErrors(ErrorFailedToFetchRecord)
	}

	item := new([]User)

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, item)
	if err != nil {
		return _handleErrors(ErrorFailedToUnmarshalRecord)
	}

	return item, nil
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(
	*User,
	error,
){
	var user User

	if _checkInvalidUserData(req, &user) {
		return _handleError(ErrorInvalidUserData)
	}

	if !validators.IsEmailValid(user.Email){
		return _handleError(ErrorInvalidEmail)
	}

	currentUser, _ := FetchUser(user.Email, tableName, dynaClient)
	if currentUser != nil && len(currentUser.Email) != 0 {
		return _handleError(ErrorUserAlreadyExists)
	}

	userAttributeValues, err := dynamodbattribute.MarshalMap(user)

	if err != nil {
		return _handleError(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item: userAttributeValues,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		return _handleError(ErrorCouldNotDynamoPutItem)
	}

	return &user, nil
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(
	*User,
	error,
){
	var user User

	if _checkInvalidUserData(req, &user) {
		return _handleError(ErrorInvalidUserData)
	}

	currentUser, _ := FetchUser(user.Email, tableName, dynaClient)
	if currentUser != nil && len (currentUser.Email) == 0 {
		return _handleError(ErrorUserDoesNotExist)
	}

	userAttributeValues, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return _handleError(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item: userAttributeValues,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		return _handleError(ErrorCouldNotDynamoPutItem)
	}

	return &user, nil
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) error{

	email := req.QueryStringParameters["email"]

	input := &dynamodb.DeleteItemInput{
		Key: _generateUserKey(email),
		TableName: aws.String(tableName),
	}

	_, err := dynaClient.DeleteItem(input)
	if err != nil {
		return errors.New(ErrorCouldNotDeleteItem)
	}

	return nil
}
