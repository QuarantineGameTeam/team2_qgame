package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

/*
	in: message to send
	out: error which possibly can happen because of wrong recipient
	or absence of rights to send the message.

	Sends given message using Client' credentials
 */
func (c *Client) SendMessage(m Message) error {
	req := url.Values{}
	req.Add("chat_id", strconv.Itoa(m.ChatID))
	req.Add("text", m.Text)

	// If message has additional InlineKeyboardMarkup
	if m.InlineMarkup.Buttons != nil {
		req.Add("reply_markup", m.InlineMarkup.stringify())
	}

	query := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?", c.token) + req.Encode()
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
