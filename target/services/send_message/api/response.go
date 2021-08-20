package api


type Response struct {
	Message string `json:"message"`
	Code    string `json:"code,omitempty"`
}

