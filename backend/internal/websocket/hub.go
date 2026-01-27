package websocket

type SessionHub struct {
	SessionID string
	Clients   map[*Client]bool
}

func NewSessionHub(id string) *SessionHub {
	return &SessionHub{
		SessionID: id,
		Clients:   make(map[*Client]bool),
	}
}
