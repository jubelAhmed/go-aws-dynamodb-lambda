package main

import (
	  "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

    "fmt"
    "log"
)

type Item struct {
	Movieid string
	Year    int
	Title   string
	Plot    string
	Rating  float64
}

func main() {
	lambda.Start(Handler)
}

func Handler() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	

	tableName := "Movies"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	

}
