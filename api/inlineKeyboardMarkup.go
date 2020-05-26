package gotg

import (
	"encoding/json"
	"log"
)

type InlineKeyboardMarkup struct {
	Buttons [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (kb InlineKeyboardMarkup) stringify() string {
	bytes, err := json.Marshal(kb)

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
