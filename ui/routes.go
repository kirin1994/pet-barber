package ui

func AddRoutes(s *Server) {
	s.Handler.HandleFunc("/companies", s.getCompanies)
	s.Handler.HandleFunc("/companies/{id}", s.getCompany)
}
