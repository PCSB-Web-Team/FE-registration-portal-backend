package db

import (
	"fmt"

	"github.com/PCSB-Web-Team/FE-registration-portal-backend/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type RegistrationsActions interface {
	CreateRegistration(*models.Registration) (*models.Registration, error)
	GetRegistration(email string) (models.Registration, error)
}

type registrationTable struct {
	tableName string
}

func createDBClient() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return dynamodb.New(sess, sess.Config.WithRegion("ap-south-1"))
}

func NewRegistrationTableInstance() RegistrationsActions {
	return &registrationTable{
		tableName: "registrations",
	}
}

func (rt *registrationTable) CreateRegistration(registration *models.Registration) (*models.Registration, error) {
	dynamodbClient := createDBClient()
	attributeValues, err := dynamodbattribute.MarshalMap(registration)
	if err != nil {
		return nil, err
	}
	tableItem := &dynamodb.PutItemInput{
		Item:      attributeValues,
		TableName: aws.String(rt.tableName),
	}
	_, err = dynamodbClient.PutItem(tableItem)
	if err != nil {
		return nil, err
	}

	return registration, nil
}

func (rt *registrationTable) GetRegistration(email string) (models.Registration, error) {
	dynamodbClient := createDBClient()
	result, err := dynamodbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(rt.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
	})

	if err != nil {
		return models.Registration{}, fmt.Errorf("error retrieving registration for '%s': %s", email, err)
	}
	if result.Item == nil {
		return models.Registration{}, fmt.Errorf("'%s' is not registered", email)
	}

	item := models.Registration{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return models.Registration{}, fmt.Errorf("failed to unmarshal record from database: %s", err)
	}
	return item, nil
}
