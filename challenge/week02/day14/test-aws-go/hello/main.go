package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func Handler(ctx context.Context, request Request) (Response, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return handleError(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	av, err := attributevalue.Marshal(request.PathParameters["id"])
	if err != nil {
		return handleError(err)
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String("MyTable"),
		Key: map[string]types.AttributeValue{
			"ID": av,
		},
	}

	resp, err := client.GetItem(ctx, input)
	if err != nil {
		return handleError(err)
	}

	item := Item{}
	err = attributevalue.UnmarshalMap(resp.Item, &item)
	if err != nil {
		return handleError(err)
	}

	responseBody, err := json.Marshal(&item)
	if err != nil {
		return handleError(err)
	}

	response := Response{
		StatusCode: http.StatusOK,
		Body:       string(responseBody),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		IsBase64Encoded: false,
	}

	return response, nil
}

func handleError(err error) (Response, error) {
	return Response{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, err
}

func main() {
	lambda.Start(Handler)
}
