package api

import (
	"context"
	"errors"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)


func getSecret(project_name string, variable_name string) (string, error) {
	path := "projects/"+project_name+"/secrets/"+variable_name

	if project_name == "" {
		return "", errors.New("Invalid project_name")
	}
	if variable_name == "" {
		return "", errors.New("Invalid token")
	}

	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("Failed to create secretmanager client: %v", err)
	}
	defer client.Close()

	// Build the request.
	req := &secretmanagerpb.GetSecretRequest{
		Name: path,
	}

	// Call the API.
	result, err := client.GetSecret(ctx, req)
	if err != nil {
		return "", fmt.Errorf("Failed to get secret: %v", err)
	}

	return result.Name, nil
}

func getJwtSigningKey(project_name string) (string, error) {
	return getSecret(project_name, "jwt_signing_key")
}

func getChatToken(project_name string, channel string) (string, error) {
	return getSecret(project_name, "discord_token_"+channel)
}

