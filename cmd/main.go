package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"monoHook/internal/cfg"
	"monoHook/internal/cpu"
	"monoHook/pkg/monobank"
	"monoHook/pkg/ynab"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var processor *cpu.CPU

func init() {
	ctx := context.Background()

	ynabToken, err := cfg.GetYnabToken(ctx)
	if err != nil {
		slog.Error("failed to get YNAB token", "error", err)
		os.Exit(1)
	}

	budgetId := os.Getenv("BUDGET_ID")
	accountId := os.Getenv("ACCOUNT_ID")
	if budgetId == "" || accountId == "" {
		slog.Error("BUDGET_ID and ACCOUNT_ID environment variables are required")
		os.Exit(1)
	}

	ynabClient := ynab.NewClient(ynabToken)
	processor = cpu.NewCPU(ynabClient, budgetId, accountId)
}

func HandleRequest(_ context.Context, event events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	if event.RequestContext.HTTP.Method == "GET" {
		return &events.APIGatewayV2HTTPResponse{StatusCode: 200}, nil
	}

	var webHookResponse monobank.WebHookResponse
	err := json.Unmarshal([]byte(event.Body), &webHookResponse)
	if err != nil {
		slog.Error("failed to unmarshal body", "error", err)
		return &events.APIGatewayV2HTTPResponse{StatusCode: 200}, nil
	}

	err = processor.AddTransaction(webHookResponse.Data.Transaction)
	if err != nil {
		slog.Error("failed to create transaction", "error", err)
	}

	return &events.APIGatewayV2HTTPResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
