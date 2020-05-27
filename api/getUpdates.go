package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) GetUpdates(offset int) []Update{
	var response Response

	req := url.Values{}
	req.Add("offset", strconv.Itoa(offset))

	query := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?", c.token) + req.Encode()
	resp, err := http.Get(query)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	return response.Result
}