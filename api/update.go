package api

/*
	Any incoming update to the bot.
 */
type Update struct{
	UpdateID int `json:"update_id"`
	Message UpdateMessage
}

func (u *Update) HasMessage() bool{
	return u.Message != UpdateMessage{}
}