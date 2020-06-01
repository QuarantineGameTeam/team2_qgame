package api

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func (c *Client) GetUploadFileRequest(chatID int, fileName string) (*http.Request, error){
	method := "/sendPhoto"
	url := "https://api.telegram.org/bot1285255270:AAFdQW1_ygN6CxQU8DzRBHLS3YLaKswLdqY" + method

	var file io.Reader
	file, err := os.Open(fileName)

	if err != nil{
		return &http.Request{}, err
	}

	buf := &bytes.Buffer{}

	w := multipart.NewWriter(buf)
	err = w.WriteField("chat_id", strconv.Itoa(chatID))

	if err != nil {
		return &http.Request{}, err
	}

	fw, err := w.CreateFormFile("photo", fileName)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = io.Copy(fw, file); err != nil {
		log.Fatal(err)
	}
	if err = w.Close(); err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		url,
		buf,
	)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	return req, nil
}

func (c *Client) SendPhoto(chatID int, fileName string) error {
	req, err := c.GetUploadFileRequest(chatID, fileName)
	if err != nil {
		return err
	}
	hc := &http.Client{}

	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK{
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(body))
	}
	return nil
}
