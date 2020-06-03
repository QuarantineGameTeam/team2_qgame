package api

/*
	Any incoming update to the bot.
*/
type Update struct {
	UpdateID      int           `json:"update_id"`
	Message       UpdateMessage `json:"message"`
	CallBackQuery CallBackQuery `json:"callback_query"`
}

func (u *Update) HasMessage() bool {
	return u.Message != UpdateMessage{}
}

func (u *Update) HasCallBackQuery() bool {
	return u.CallBackQuery != CallBackQuery{}
}
