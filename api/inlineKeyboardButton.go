package gotg

type InlineKeyboardButton struct {
	Text string `json:"text"`
	URL  string `json:"url"`
	Callback string `json:"callback_data"`
}
