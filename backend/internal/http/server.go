package httpapi

import (
	"database/sql"
	"net/http"

	"github.com/bhawani-prajapat2006/0Xnet/backend/internal/websocket"
)

type Server struct {
	db       *sql.DB
	deviceID string
}

func NewServer(db *sql.DB, deviceID string) *Server {
	return &Server{db: db, deviceID: deviceID}
}

func (s *Server) Start() {
	http.HandleFunc("/session/create", s.createSession)
	http.HandleFunc("/session/list", s.listSessions)
	http.HandleFunc("/ws", websocket.ServeWS)
	http.ListenAndServe(":8080", nil)
}
