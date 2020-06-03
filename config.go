package main

const (
	botToken = "1285255270:AAFdQW1_ygN6CxQU8DzRBHLS3YLaKswLdqY"

	//User state constants
	StateAFK          = iota //passive state for user
	StateChangingName = iota //user is changing nickname
	StateWriting      = iota //user answers the question
	StateWaiting      = iota //user is waiting for his turn
	StateThinking     = iota //there is user's move now
)
