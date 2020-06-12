package config

const (
	BotToken = "0000000000:CorrectToken"
	DevID    = 662834330
)

// Separated in order to make iotas work properly
const (
	//User state constants
	StateNone         = iota // no interaction possible at all
	StateAFK          = iota // passive state for user
	StateChangingName = iota // user is changing nickname
	StateWriting      = iota // user answers the question
	StateWaiting      = iota // user is waiting for his turn
	StateThinking     = iota // there is user's move now
)
