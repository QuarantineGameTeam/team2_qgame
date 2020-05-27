package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	token string
}

func (c *Client) SetToken(token string) error {
	c.token = token

	query := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", token)

	resp, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	response := Response{}
	err = json.Unmarshal(bytes, &response)

	if response.OK {
		return nil
	}
	return errors.New(string(bytes))
}