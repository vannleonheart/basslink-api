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

	userEndpoint := mainEndpoint.Group("/users", s.shouldBeUser([]string{}))
	userEndpoint.Get("/", s.handleGetUsers)
	userEndpoint.Get("/:id", s.handleGetUser)

	contactEndpoint := mainEndpoint.Group("/contacts", s.shouldBeUser([]string{}))
	contactEndpoint.Get("/", s.handleGetContacts)
	contactEndpoint.Get("/:id", s.handleGetContact)

	contactEndpoint.Get("/:contactId/accounts", s.handleGetContactAccounts)

	disbursementEndpoint := mainEndpoint.Group("/disbursements", s.shouldBeUser([]string{}))
	disbursementEndpoint.Get("/", s.handleGetDisbursements)
	disbursementEndpoint.Get("/:id", s.handleGetDisbursement)
}
