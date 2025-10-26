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

	senderEndpoint := mainEndpoint.Group("/senders", s.shouldBeUser([]string{}))
	senderEndpoint.Get("/", s.handleGetUsers)
	senderEndpoint.Get("/:id", s.handleGetUser)
	senderEndpoint.Post("/", s.handleCreateUser)
	senderEndpoint.Put("/:id", s.handleUpdateUser)

	recipientEndpoint := mainEndpoint.Group("/recipients", s.shouldBeUser([]string{}))
	recipientEndpoint.Get("/", s.handleGetRecipients)
	recipientEndpoint.Get("/:id", s.handleGetRecipient)
	recipientEndpoint.Post("/", s.handleCreateRecipient)
	recipientEndpoint.Put("/:id", s.handleUpdateRecipient)
	recipientEndpoint.Delete("/:id", s.handleDeleteRecipient)

	recipientEndpoint.Post("/:recipientId/documents", s.handleCreateRecipientDocument)
	recipientEndpoint.Put("/:recipientId/documents/:documentId", s.handleUpdateRecipientDocument)
	recipientEndpoint.Delete("/:recipientId/documents/:documentId", s.handleDeleteRecipientDocument)

	remittanceEndpoint := mainEndpoint.Group("/remittances", s.shouldBeUser([]string{}))
	remittanceEndpoint.Get("/", s.handleGetRemittances)
	remittanceEndpoint.Get("/:id", s.handleGetRemittance)
	remittanceEndpoint.Post("/", s.handleCreateRemittance)
	remittanceEndpoint.Post("/:id/complete", s.handleCompleteRemittance)

	submissionEndpoint := mainEndpoint.Group("/submissions", s.shouldBeUser([]string{}))
	submissionEndpoint.Get("/", s.handleGetSubmissions)
	submissionEndpoint.Post("/:id/accept", s.handleAcceptSubmission)
	submissionEndpoint.Post("/:id/reject", s.handleRejectSubmission)

	appointmentEndpoint := mainEndpoint.Group("/appointments", s.shouldBeUser([]string{}))
	appointmentEndpoint.Get("/", s.handleGetAppointments)
}
