package secrets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

const (
	// VersionStage defaults to AWSCURRENT if unspecified
	SecretManagerVersionStageDefault = "AWSCURRENT"
)

// GetSecrets retrieves the secret value from AWS Secrets Manager.
// The secret value is returned as a string.
//
// About the GetSecretValue API, see
//
//	For a list of exceptions thrown, see
//	https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
func GetSecrets(region, key string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return "", err
	}

	svc := secretsmanager.NewFromConfig(cfg)

	result, err := svc.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(key),
		VersionStage: aws.String(SecretManagerVersionStageDefault),
	})
	if err != nil {
		return "", err
	}

	return *result.SecretString, nil
}
