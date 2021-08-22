package api

import (
	"context"
	"errors"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func getSecret(project_id string, secret_name string) (string, error) {
	secret_id := "projects/" + project_id + "/secrets/" + secret_name + "/versions/latest"

	if project_id == "" {
		return "", errors.New("Invalid project_id")
	}
	if secret_id == "" {
		return "", errors.New("Invalid secret name")
	}

	// Create the client
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("Failed to create secretmanager client: %v", err)
	}
	defer client.Close()

	// Build the request
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secret_id,
	}

	// Call the API then check the result
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("Failed to get secret: %v", err)
	}
	if result.Payload == nil {
		return "", errors.New("No payload found")
	}
	value := string(result.Payload.Data)
	if value == "" {
		return "", errors.New("Empty secret value")
	}

	return value, nil
}

func getJwtSigningKey(project_id string) (string, error) {
	return getSecret(project_id, "jwt_signing_key")
}

func getChatToken(project_id string, channel string) (string, error) {
	return getSecret(project_id, "discord_token_"+channel)
}
