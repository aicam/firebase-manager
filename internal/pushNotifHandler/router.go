package pushNotifHandler

func (s *Server) Routes() {
	// s.Router.HandleFunc("/login/{username}/{password}", s.LoginOnce(db))
	s.Router.HandleFunc("/add_user/{username}/{score}", s.addUser())
	s.Router.HandleFunc("/set_token/{username}/{token}", s.setToken())
}
