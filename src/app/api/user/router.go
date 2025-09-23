package user

func (s *Service) InitRouter() {
	mainEndpoint := s.App.HttpServer.Handler.Group("/user", s.handleSession)

	authEndpoint := mainEndpoint.Group("/auth", s.shouldBeGuest)
	authEndpoint.Post("/signin", s.handleSignIn)

	contactEndpoint := mainEndpoint.Group("/contacts", s.shouldBeUser)
	contactEndpoint.Get("/", s.handleGetContacts)
	contactEndpoint.Get("/:id", s.handleGetContact)
	contactEndpoint.Post("/", s.handleCreateContact)
	contactEndpoint.Put("/:id", s.handleUpdateContact)
	contactEndpoint.Delete("/:id", s.handleDeleteContact)

	contactEndpoint.Post("/:contactId/documents", s.handleCreateContactDocument)
	contactEndpoint.Put("/:contactId/documents/:documentId", s.handleUpdateContactDocument)
	contactEndpoint.Delete("/:contactId/documents/:documentId", s.handleDeleteContactDocument)

	contactEndpoint.Post("/:contactId/accounts", s.handleCreateContactAccount)
	contactEndpoint.Put("/:contactId/accounts/:accountId", s.handleUpdateContactAccount)
	contactEndpoint.Delete("/:contactId/accounts/:accountId", s.handleDeleteContactAccount)

	disbursementEndpoint := mainEndpoint.Group("/disbursements", s.shouldBeUser)
	disbursementEndpoint.Get("/", s.handleGetDisbursements)
	disbursementEndpoint.Get("/:id", s.handleGetDisbursement)
}
