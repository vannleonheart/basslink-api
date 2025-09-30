package common

func (s *Service) InitRouter() {
	s.App.HttpServer.Handler.Get("/currencies", s.GetCurrenciesHandler)
	s.App.HttpServer.Handler.Post("/upload", s.handleSession, s.shouldBeUser, s.FileUploadHandler)
}
