package api

type CallBackQuery struct {
	ID string `json:"id"`
	FromUser User `json:"from"`
	Message UpdateMessage `json:"message"`
	CallBackData string `json:"data"`
}