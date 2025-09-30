package common

import (
	"CRM/src/lib/basslink"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) FileUploadHandler(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	path := "uploads/general"

	if ctxAdmin := c.Locals("admin"); ctxAdmin != nil {
		admin := ctxAdmin.(*basslink.Administrator)
		path = "uploads/admin/" + admin.Id
	}

	if ctxAgent := c.Locals("agent"); ctxAgent != nil {
		agent := ctxAgent.(*basslink.AgentUser)
		path = "uploads/agent/" + agent.Agent.Id
	}

	if ctxUser := c.Locals("user"); ctxUser != nil {
		user := ctxUser.(*basslink.User)
		path = "uploads/user/" + user.Id
	}

	result, apErr := s.handleUploadFile(path, form)
	if apErr != nil {
		return apErr
	}

	return basslink.NewSuccessResponse(c, "FILE_UPLOAD_SUCCESS", result)
}

func (s *Service) GetCurrenciesHandler(c *fiber.Ctx) error {
	result, apErr := s.handleGetCurrencies()
	if apErr != nil {
		return apErr
	}

	return basslink.NewSuccessResponse(c, "CURRENCIES_LIST_SUCCESS", result)
}
