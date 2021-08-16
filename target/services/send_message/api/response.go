package api

type Response struct {
	message string `json:"message,omitempty"`
	code    string `json:"code,omitempty"`
}
