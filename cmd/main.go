package main

import (
	"context"
	"encoding/json"
	"fmt"
	"monoHook/internal/cfg"
	"monoHook/internal/cpu"
	"monoHook/pkg/monobank"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/brunomvsouza/ynab.go"
)

func HandleRequest(ctx context.Context, event events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	if event.RequestContext.HTTP.Method == "GET" {
		return &events.APIGatewayV2HTTPResponse{StatusCode: 200}, nil
	}

	ynabToken := cfg.GetYnabToken(ctx)

	fmt.Printf("Event: %v\n", event.Body)

	var webHookResponse monobank.WebHookResponse
	err := json.Unmarshal([]byte(event.Body), &webHookResponse)
	if err != nil {
		fmt.Printf("Couldn't unmarshal body, %v\n", err)
		answer := events.APIGatewayV2HTTPResponse{StatusCode: 400, Body: "Invalid format"}
		return &answer, err
	}

	budgetId := os.Getenv("BUDGET_ID")
	accountId := os.Getenv("ACCOUNT_ID")

	ynabClient := ynab.NewClient(ynabToken)

	c := cpu.NewCPU(ynabClient, budgetId, accountId)
	err = c.AddTransaction(webHookResponse.Data.Transaction)
	if err != nil {
		fmt.Printf("Couldn't create transaction, %v\n", err)
		answer := events.APIGatewayV2HTTPResponse{StatusCode: 500, Body: "Couldn't create transaction"}
		return &answer, err
	}

	return &events.APIGatewayV2HTTPResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
