package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
func (c *Client) SendMessage(m Message) (UpdateMessage, error) {
	method := "/sendMessage"

	req := url.Values{}
	req.Add("chat_id", strconv.Itoa(m.ChatID))
	req.Add("text", m.Text)

	// If message has additional InlineKeyboardMarkup
	if m.InlineMarkup.Buttons != nil {
		req.Add("reply_markup", m.InlineMarkup.stringify())
	}

	query := fmt.Sprintf("%s%s%s?", botEntry, c.token, method) + req.Encode()
	resp, err := http.Get(query)

	if err != nil {
		return UpdateMessage{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return UpdateMessage{}, err
	}

	if resp.StatusCode == 200 {
		var response struct{
			OK bool `json:"ok"`
			Result struct {
				UMessage UpdateMessage
			}
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return UpdateMessage{}, err
		}
		if !response.OK {
			return UpdateMessage{}, err
		}

		return response.Result.UMessage, nil
	}

	return UpdateMessage{}, errors.New(string(body))
}
