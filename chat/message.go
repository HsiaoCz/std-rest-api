package chat

type Message struct {
	ClientID string
	Text     string
}

type WSMessage struct {
	Text    string `json:"text"`
	Headers any    `json:"HEADERS"`
}
