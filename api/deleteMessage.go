package api

import (
	"fmt"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"log"
)

func (c *Client) DeleteMessage(message UpdateMessage) error {
	req := url.Values{}
	req.Add("chat_id", strconv.Itoa(message.Chat.ID))
	req.Add("message_id", strconv.Itoa(message.MessageID))

	method := "/deleteMessage"
	query := fmt.Sprintf("%s%s%s?", botEntry, c.token, method) + req.Encode()
	resp, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		return nil
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return errors.New(string(body))
}