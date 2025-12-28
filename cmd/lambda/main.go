package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	"coinbase-automation/internal/config"
	"coinbase-automation/internal/handler"
	"coinbase-automation/internal/secret"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	secretClient, err := secret.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create secret client: %v", err)
	}

	s, err := secretClient.Fetch(ctx, cfg.SecretName)
	if err != nil {
		log.Fatalf("Failed to fetch secret: %v", err)
	}

	h := handler.New()

	lambda.Start(h.Run)
}
