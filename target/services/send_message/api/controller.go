package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	echo "github.com/labstack/echo/v4"
)


/*
 * Controller functions
 */
func isAlive(c echo.Context) error {
	response := Response{Message: "Hello I am still alive !"}
	return c.JSON(http.StatusOK, response)
}

func sendMessageFunc(project_name string) (func (c echo.Context) error) {
	sendMessage := func(c echo.Context) error {
		// Extract the channel_id from the URL
		channel_id := c.Param("channel")

		// Check if the JWT token allows this channel
		jwt := c.QueryParam("jwt_key")
		signingKey, err := getJwtSigningKey(project_name)
		if err != nil {
			response := Response{Message: "Error in secret loading"}
			c.JSON(http.StatusInternalServerError, response)
			return err
		}

		if !hasChannel(channel_id, jwt, signingKey) {
			message := "Unauthorized channel"
			response := Response{Message: message}
			c.JSON(http.StatusUnauthorized, response)
			return errors.New(message)
		}

		// Extract all the required params to post the message
		token, err := getChatToken(project_name, channel_id)
		if err != nil {
			response := Response{Message: "Error in secret loading"}
			c.JSON(http.StatusInternalServerError, response)
			return err
		}

		url := getHTTPurl(channel_id, token)
		username := extractUsername(jwt, signingKey)

		// Get the content from the HTTP body
		defer c.Request().Body.Close()
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			response := Response{Message: "Error in body reading"}
			c.JSON(http.StatusInternalServerError, response)
			return err
		}

		request_body := Request{}
		json_string := string(body)
		err = json.Unmarshal([]byte(json_string), &request_body)
		if err != nil {
			response := Response{Message: "Error in content deserialization"}
			c.JSON(http.StatusInternalServerError, response)
			return err
		}
		content := request_body.Content

		// Call the API
		httpCode, message, err := callAPI(url, username, content)
		response := Response{Message: message}

		c.JSON(httpCode, response)
		return err
	}
	return sendMessage
}

