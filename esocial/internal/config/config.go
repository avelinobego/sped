package config

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func init() {

	configureLogger()

	if err := loadSecretsToEnv("postgre-secret", "us-east-1"); err != nil {
		slog.Error("Failed to load secrets", "erro", err)
	}
}

func configureLogger() {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	instance := slog.New(handler)

	// Set as the default logger for the slog package
	slog.SetDefault(instance)
}

func loadSecretsToEnv(secretName string, region string) error {
	// 1. Load AWS Configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return err
	}

	// 2. Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(cfg)

	// 3. Get the secret value
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return err
	}

	// 4. Secrets Manager returns a string, usually JSON formatted
	var secretMap map[string]string
	err = json.Unmarshal([]byte(*result.SecretString), &secretMap)
	if err != nil {
		return err
	}

	// 5. Loop through the map and set environment variables
	for key, value := range secretMap {
		os.Setenv(key, value)
		slog.Info("Loaded secret into ENV", "chave", key)
	}

	return nil
}
