package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type jsonbody struct {
	Content string `json:"content"`
}

/*
 * Controllers of the routes
 */
func isAlive(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello I am still alive !")
}

func sendMessage(c echo.Context) error {
	// Extract the token from the secret vault
	channel_id := c.Param("channel")
	token := getToken(channel_id)

	// Extract all the required params to post the message
	url := getHTTPurl(channel_id, token)
	username := getJWTusername(token)

	defer c.Request().Body.Close()
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}
	json_data := string(body)

	body_elmt := jsonbody{}
	err = json.Unmarshal([]byte(json_data), &body_elmt)
	if err != nil {
		return err
	}
	content := body_elmt.Content

	return postMessage(url, username, content)
}
