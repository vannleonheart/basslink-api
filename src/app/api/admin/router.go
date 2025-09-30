package admin

func (s *Service) InitRouter() {
	mainEndpoint := s.App.HttpServer.Handler.Group("/admin", s.handleSession)

	authEndpoint := mainEndpoint.Group("/auth", s.shouldBeGuest)
	authEndpoint.Post("/signin", s.handleSignIn)

	accountEndpoint := mainEndpoint.Group("/account", s.shouldBeUser([]string{}))
	accountEndpoint.Get("/profile", s.handleGetProfile)
	accountEndpoint.Patch("/password", s.handleUpdatePassword)

	userEndpoint := mainEndpoint.Group("/users", s.shouldBeUser([]string{}))
	userEndpoint.Get("/", s.handleGetUsers)
	userEndpoint.Get("/:id", s.handleGetUser)
	userEndpoint.Post("/", s.handleCreateUser)
	userEndpoint.Put("/:id", s.handleUpdateUser)
	userEndpoint.Patch("/:id/status", s.handleToggleUserEnable)
	userEndpoint.Delete("/:id", s.handleDeleteUser)

	agentEndpoint := mainEndpoint.Group("/agents", s.shouldBeUser([]string{}))
	agentEndpoint.Get("/", s.handleGetAgents)
	agentEndpoint.Get("/:id", s.handleGetAgent)
	agentEndpoint.Post("/", s.handleCreateAgent)
	agentEndpoint.Put("/:id", s.handleUpdateAgent)
	agentEndpoint.Patch("/:id/status", s.handleToggleAgentEnable)

	disbursementEndpoint := mainEndpoint.Group("/disbursements", s.shouldBeUser([]string{}))
	disbursementEndpoint.Get("/", s.handleGetDisbursements)
	disbursementEndpoint.Get("/:id", s.handleGetDisbursement)
}
