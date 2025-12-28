package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	"coinbase-automation/internal/config"
	"coinbase-automation/internal/handler"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	h := handler.New()

	lambda.Start(h.Run)
}
