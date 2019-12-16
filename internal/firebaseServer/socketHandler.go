package firebaseServer

import (
	"github.com/gin-gonic/gin"
	"log"
)

func (s *Server) handleSearchUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := s.SocketConnection.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Print(err)
			return
		}
		conn.SetReadLimit(30)
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err, "reading")
				return
			}
			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}
	}
}
