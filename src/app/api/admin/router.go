package admin

func (s *Service) InitRouter() {
	mainEndpoint := s.App.HttpServer.Handler.Group("/admin", s.handleSession)

	authEndpoint := mainEndpoint.Group("/auth", s.shouldBeGuest)
	authEndpoint.Post("/signin", s.handleSignIn)

	accountEndpoint := mainEndpoint.Group("/account", s.shouldBeUser([]string{}))
	accountEndpoint.Get("/profile", s.handleGetProfile)
	accountEndpoint.Patch("/password", s.handleUpdatePassword)

	adminUserEndpoint := mainEndpoint.Group("/admin_users", s.shouldBeUser([]string{}))
	adminUserEndpoint.Get("/", s.handleGetAdminUsers)
	adminUserEndpoint.Get("/:id", s.handleGetAdminUser)
	adminUserEndpoint.Post("/", s.handleCreateAdminUser)
	adminUserEndpoint.Put("/:id", s.handleUpdateAdminUser)
	adminUserEndpoint.Patch("/:id/status", s.handleToggleAdminUserEnable)
	adminUserEndpoint.Delete("/:id", s.handleDeleteAdminUser)

	agentEndpoint := mainEndpoint.Group("/agents", s.shouldBeUser([]string{}))
	agentEndpoint.Get("/", s.handleGetAgents)
	agentEndpoint.Get("/:id", s.handleGetAgent)
	agentEndpoint.Post("/", s.handleCreateAgent)
	agentEndpoint.Put("/:id", s.handleUpdateAgent)
	agentEndpoint.Patch("/:id/status", s.handleToggleAgentEnable)

	userEndpoint := mainEndpoint.Group("/senders", s.shouldBeUser([]string{}))
	userEndpoint.Get("/", s.handleGetUsers)
	userEndpoint.Get("/:id", s.handleGetUser)

	recipientEndpoint := mainEndpoint.Group("/recipients", s.shouldBeUser([]string{}))
	recipientEndpoint.Get("/", s.handleGetRecipients)
	recipientEndpoint.Get("/:id", s.handleGetRecipient)

	remittanceEndpoint := mainEndpoint.Group("/remittances", s.shouldBeUser([]string{}))
	remittanceEndpoint.Get("/", s.handleGetRemittances)
	remittanceEndpoint.Get("/:id", s.handleGetRemittance)
}
