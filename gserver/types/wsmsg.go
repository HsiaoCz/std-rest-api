package types

type WsMessage struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}
