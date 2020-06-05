package api

/*
	Message to send only.
	Cannot be used to parse incoming messages.
	For that, please refer to UpdateMessage.
 */
type Message struct {
	ChatID       int
	Text         string
	InlineMarkup InlineKeyboardMarkup
}
