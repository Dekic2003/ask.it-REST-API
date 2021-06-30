package utils

type ErrMsg struct {
	Message string `json:"message"`
	Error error `json:"error"`
}
type SuccessMsg struct {
	Data interface{} `json:"data"`
	Success bool `json:"success"`
}