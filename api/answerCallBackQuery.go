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

func (c *Client) AnswerCallBackQuery(query CallBackQuery, text string, alert bool) error {
	method := "/answerCallbackQuery"
	values := url.Values{}

	values.Add("callback_query_id", query.ID)
	if text != "" {
		values.Add("text", text)
	}
	values.Add("show_alert", strconv.FormatBool(alert))
	req := fmt.Sprintf("%s%s%s?", botEntry, c.token, method) + values.Encode()

	res, err := http.Get(req)
	if err != nil {
		return err
	}

	var parsedResponse struct {
		OK          bool   `json:"ok"`
		ErrCode     int    `json:"error_code"`
		Description string `json:"description"`
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &parsedResponse)
	if err != nil {
		return err
	}

	if parsedResponse.OK == false {
		return errors.New(fmt.Sprintf("STATUSCODE %d : %s", parsedResponse.ErrCode, parsedResponse.Description))
	}
	return nil
}
