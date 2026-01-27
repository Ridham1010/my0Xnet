package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/bhawani-prajapat2006/0Xnet/backend/internal/service"
)

func (s *Server) createSession(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name string `json:"name"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	session, err := service.CreateSession(s.db, body.Name, s.deviceID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(session)
}

func (s *Server) listSessions(w http.ResponseWriter, r *http.Request) {
	sessions, _ := service.ListSessions(s.db)
	json.NewEncoder(w).Encode(sessions)
}
