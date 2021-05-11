package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("You sent me, %v", request.Body)
	fmt.Println("With headers:")
	for key, value := range request.Headers {
		fmt.Printf("%s: %s\n", key, value)
	}

	name := request.QueryStringParameters["name"]
	if name == "" {
		name = "Friend"
	}
	return events.APIGatewayProxyResponse{Body: fmt.Sprintf("Hello, %s!", name), StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
