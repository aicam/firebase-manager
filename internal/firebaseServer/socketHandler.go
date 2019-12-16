package firebaseServer

import "net/http"

func (s *Server) handleSearchUsername() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		conn, err := s.SocketConnection.Upgrade(writer, request, nil)
		if err != nil {

		}
	}
}
