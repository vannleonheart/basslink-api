package agent

func (s *Service) InitRouter() {
	mainEndpoint := s.App.HttpServer.Handler.Group("/agent", s.handleSession)

	authEndpoint := mainEndpoint.Group("/auth", s.shouldBeGuest)
	authEndpoint.Post("/signin", s.handleSignIn)

	accountEndpoint := mainEndpoint.Group("/account", s.shouldBeUser([]string{}))
	accountEndpoint.Get("/profile", s.handleGetProfile)
	accountEndpoint.Patch("/password", s.handleUpdatePassword)

	agentUserEndpoint := mainEndpoint.Group("/agent_users", s.shouldBeUser([]string{}))
	agentUserEndpoint.Get("/", s.handleGetAgentUsers)
	agentUserEndpoint.Get("/:id", s.handleGetAgentUser)
	agentUserEndpoint.Post("/", s.handleCreateAgentUser)
	agentUserEndpoint.Put("/:id", s.handleUpdateAgentUser)
	agentUserEndpoint.Patch("/:id/status", s.handleToggleAgentUserEnable)
	agentUserEndpoint.Delete("/:id", s.handleDeleteAgentUser)

	userEndpoint := mainEndpoint.Group("/users", s.shouldBeUser([]string{}))
	userEndpoint.Get("/", s.handleGetUsers)
	userEndpoint.Get("/:id", s.handleGetUser)
	userEndpoint.Post("/", s.handleCreateUser)
	userEndpoint.Put("/:id", s.handleUpdateUser)
	userEndpoint.Patch("/:id/status", s.handleToggleUserEnable)

	contactEndpoint := mainEndpoint.Group("/contacts", s.shouldBeUser([]string{}))
	contactEndpoint.Get("/", s.handleGetContacts)
	contactEndpoint.Get("/:id", s.handleGetContact)
	contactEndpoint.Post("/", s.handleCreateContact)
	contactEndpoint.Put("/:id", s.handleUpdateContact)
	contactEndpoint.Delete("/:id", s.handleDeleteContact)

	contactEndpoint.Post("/:contactId/documents", s.handleCreateContactDocument)
	contactEndpoint.Put("/:contactId/documents/:documentId", s.handleUpdateContactDocument)
	contactEndpoint.Delete("/:contactId/documents/:documentId", s.handleDeleteContactDocument)

	contactEndpoint.Get("/:contactId/accounts", s.handleGetContactAccounts)
	contactEndpoint.Post("/:contactId/accounts", s.handleCreateContactAccount)
	contactEndpoint.Put("/:contactId/accounts/:accountId", s.handleUpdateContactAccount)
	contactEndpoint.Delete("/:contactId/accounts/:accountId", s.handleDeleteContactAccount)

	disbursementEndpoint := mainEndpoint.Group("/disbursements", s.shouldBeUser([]string{}))
	disbursementEndpoint.Get("/", s.handleGetDisbursements)
	disbursementEndpoint.Get("/:id", s.handleGetDisbursement)
	disbursementEndpoint.Post("/", s.handleCreateDisbursement)
}
