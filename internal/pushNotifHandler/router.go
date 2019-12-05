package pushNotifHandler

func (s *Server) Routes() {
	// s.Router.HandleFunc("/login/{username}/{password}", s.LoginOnce(db))
	s.Router.GET("/add_user/:username/:score", s.addUser())
	s.Router.GET("/set_token/:username/:token", s.setToken())
	s.Router.POST("/send_notif/:username/:title/:topic", s.sendNotification())
}
