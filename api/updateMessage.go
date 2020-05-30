package api

/*
	Stands for the incoming message in update.
	Cannot be used as message to send.
	For that, please refer to Message.
 */
type UpdateMessage struct {
	MessageID int `json:"message_id"`
	User User `json:"from"`
	Date int64 `json:"date"`
	Text string `json:"text"`
}