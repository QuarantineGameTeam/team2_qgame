package api

type UpdateMessage struct {
	MessageID int `json:"message_id"`
	User User `json:"from"`
	Date int64 `json:"date"`
	Text string `json:"text"`
}