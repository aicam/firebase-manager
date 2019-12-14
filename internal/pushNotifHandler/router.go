package pushNotifHandler

func (s *Server) Routes() {
	s.Router.GET("/add_user/:username/:score", s.addUser())
	s.Router.GET("/set_token/:username/:token", s.setToken())
	s.Router.POST("/send_notif/:username/:title", s.sendNotification())
	s.Router.GET("/add_single_score/:username/:score", s.addScore())
	s.Router.POST("/add_multiple_score", s.addMultipleScore())
	s.Router.POST("/send_multiple_notification", s.sendMultipleNotification())
	s.Router.GET("/get_users/:offset/:limit", s.getUsers())
	s.Router.POST("/get_failed_messages", s.getFailedMessagesByDate())
}
