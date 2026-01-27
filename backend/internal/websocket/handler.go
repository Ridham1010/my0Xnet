package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	DeviceID string
	Conn     *websocket.Conn
	Session  string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	var msg map[string]string
	conn.ReadJSON(&msg)

	// First message must be join-session
	if msg["type"] != "join-session" {
		conn.Close()
		return
	}

	// Approval logic handled by HTTP / DB
	// If approved → keep connection
	for {
		var incoming map[string]interface{}

		err := conn.ReadJSON(&incoming)
		if err != nil {
			break
		}

		// process incoming data here

		err = conn.WriteJSON(incoming)
		if err != nil {
			break
		}
	}
}
