package chat

type Hub struct {
	clients    map[*Client]bool
	brodcast   chan *Message
	register   chan *Client
	unregister chan *Client
}
