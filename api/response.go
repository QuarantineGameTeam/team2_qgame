package api

type Response struct {
	OK bool `json:"ok"`

	Result []Update `json:"result"`

	ErrorCode int `json:"error_code"`
	Description string `json:"description"`
}