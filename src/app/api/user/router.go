package user

func (s *Service) InitRouter() {
	mainEndpoint := s.App.HttpServer.Handler.Group("/user", s.handleSession)

	authEndpoint := mainEndpoint.Group("/auth", s.shouldBeGuest)
	authEndpoint.Post("/signin", s.handleSignIn)

	accountEndpoint := mainEndpoint.Group("/account", s.shouldBeUser)
	accountEndpoint.Get("/profile", s.handleGetProfile)
	accountEndpoint.Patch("/password", s.handleUpdatePassword)

	disbursementEndpoint := mainEndpoint.Group("/disbursements", s.shouldBeUser)
	disbursementEndpoint.Get("/", s.handleGetDisbursements)
	disbursementEndpoint.Get("/:id", s.handleGetDisbursement)
}
