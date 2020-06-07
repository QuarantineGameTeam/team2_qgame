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

func (c *Client) DeleteMessage(UpdateMessage) error{
	req := url.Values{}
	req.Add("chat_id", strconv.Itoa(message.FromUser.ID))
	req.Add("message_id", message.MessageID)

	method := "/deleteMessage"
	query := fmt.Sprintf("https://api.telegram.org/bot%s/deleteMessage?", c.token) + req.Encode() + method
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