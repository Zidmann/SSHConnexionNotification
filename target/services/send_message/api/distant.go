package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

/*
 * Functions to interact with external services
 */

// Build the final URL
func getHTTPurl(channel_id string, token string) string {
	return "https://discord.com/api/webhooks/" + channel_id + "/" + token
}

// Send the message
func callAPI(url string, username string, content string) (int, string, error) {
	request := Request{Username: username, Content: content}
	reqBody, err := json.Marshal(request)

	if err != nil {
		return http.StatusInternalServerError, "Error in distant body serialization", err
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return http.StatusBadGateway, "Error during the POST request", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, "Error in distant body reading", err
	}
	json_string := string(body)

	return http.StatusOK, json_string, nil
}
