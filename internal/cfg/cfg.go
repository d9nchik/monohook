package cfg

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type YnabToken struct {
	Token string `json:"token"`
}

func GetYnabToken(ctx context.Context) (string, error) {
	secretName := "ynabToken"
	region := "eu-central-1"

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return "", fmt.Errorf("failed to load AWS config: %w", err)
	}

	svc := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to get secret value: %w", err)
	}

	var ynabToken YnabToken
	err = json.Unmarshal([]byte(*result.SecretString), &ynabToken)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal secret: %w", err)
	}

	return ynabToken.Token, nil
}
