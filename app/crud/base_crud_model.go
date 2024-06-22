package crud

// Base model for responses
type BaseResponse struct {
	Error   bool   `json:"error"`
	Code    string `json:"code"`
	Message string `json:"msg"`
}
