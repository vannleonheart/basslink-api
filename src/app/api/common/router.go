package common

func (s *Service) InitRouter() {
	s.App.HttpServer.Handler.Get("/currencies", s.GetCurrenciesHandler)
	s.App.HttpServer.Handler.Post("/rates", s.GetRateHandler)
	s.App.HttpServer.Handler.Post("/upload", s.handleSession, s.shouldBeUser, s.FileUploadHandler)
	s.App.HttpServer.Handler.Post("/upload-public", s.PublicFileUploadHandler)
	s.App.HttpServer.Handler.Post("/appointments", s.CreateAppointmentHandler)
	s.App.HttpServer.Handler.Post("/transactions/search", s.SearchTransactionHandler)
	s.App.HttpServer.Handler.Post("/transactions/create", s.CreateTransactionHandler)
}
