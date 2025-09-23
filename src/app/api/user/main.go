package user

import "CRM/src/lib/basslink"

type Service struct {
	App *basslink.App
}

func New(app *basslink.App) *Service {
	return &Service{
		App: app,
	}
}
