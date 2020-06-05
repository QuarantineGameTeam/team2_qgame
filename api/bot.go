package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	botEntry     = "https://api.telegram.org/bot"
	testToken    = "1285255270:AAFdQW1_ygN6CxQU8DzRBHLS3YLaKswLdqY"
	testSendToID = 662834330
)

type Client struct {
	token string
}

func (c *Client) SetToken(token string) error {
	c.token = token
	method := "/getMe"

	query := fmt.Sprintf("%s%s%s?", botEntry, token, method)

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

func NewClient(token string) (c *Client, err error) {
	c = new(Client)
	err = c.SetToken(token)
	if err != nil {
		return &Client{}, err
	}
	return
}
