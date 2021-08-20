package api


type Request struct {
	Username string `json:"username,omitempty"`
	Content  string `json:"content"`
}

