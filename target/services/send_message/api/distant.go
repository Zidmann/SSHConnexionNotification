package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getHTTPurl(channel_id string, token string) string {
	return "https://discord.com/api/webhooks/" + channel_id + "/" + token
}

/*
 * Functions to interact with different services
 */
func postMessage(url string, username string, content string) error {
	reqBody, err := json.Marshal(map[string]string{
		"content":  content,
		"username": username,
	})

	if err != nil {
		print(err)
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
		return err
	}
	fmt.Println(string(body))
	return nil
}
