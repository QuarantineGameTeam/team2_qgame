package gotg

type Message struct {
	ChatID       int
	Text         string
	InlineMarkup InlineKeyboardMarkup
}
