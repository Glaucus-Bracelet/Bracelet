package password

type Response struct {
	Error string `json:"error,omitempty"`
	Token string `json:"token,omitempty"`
}
