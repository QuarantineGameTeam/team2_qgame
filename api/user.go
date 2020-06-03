package api

type User struct {
	ID        int    `json:"id"`
	State     int    `json:"state"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	LangCode  string `json:"language_code"`
}
