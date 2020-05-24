package api

type Update struct{
	UpdateID int `json:"update_id"`
	Message UpdateMessage
}
