package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("critical error: %v", err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	name := request.QueryStringParameters["name"]
	if name == "" {
		name = "Friend"
	}
	table := "goexampletable"
	resp, err := svc.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &table,
		Item: map[string]types.AttributeValue{
			"RequestId": &types.AttributeValueMemberS{Value: uuid.NewString()},
			"Name":      &types.AttributeValueMemberS{Value: name},
		},
	})
	if err != nil {
		log.Fatalf("critical error: %v", err)
	}
	log.Println(resp)

	fmt.Printf("You sent me, %v", request.Body)
	fmt.Println("With headers:")
	for key, value := range request.Headers {
		fmt.Printf("%s: %s\n", key, value)
	}

	return events.APIGatewayProxyResponse{Body: fmt.Sprintf("Hello, %s!", name), StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
