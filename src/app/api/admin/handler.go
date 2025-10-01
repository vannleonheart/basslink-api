package admin

import (
	"CRM/src/lib/basslink"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) handleSignIn(c *fiber.Ctx) error {
	var req SignInRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	result, err := s.signIn(&req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "AUTH_SIGNIN_SUCCESS", result)
}

func (s *Service) handleGetProfile(c *fiber.Ctx) error {
	adminUser := c.Locals("admin").(*basslink.Administrator)
	result, err := s.getProfile(adminUser)
	if err != nil {
		return err
	}
	return basslink.NewSuccessResponse(c, "ACCOUNT_PROFILE_GET_SUCCESS", result)
}

func (s *Service) handleUpdatePassword(c *fiber.Ctx) error {
	var req UpdatePasswordRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	adminUser := c.Locals("admin").(*basslink.Administrator)
	err := s.updatePassword(adminUser, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "ACCOUNT_PASSWORD_UPDATE_SUCCESS", nil)
}

func (s *Service) handleGetAdminUsers(c *fiber.Ctx) error {
	result, err := s.getAdminUsers()
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_LIST_SUCCESS", result)
}

func (s *Service) handleGetAdminUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	result, err := s.getAdminUser(userId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_GET_SUCCESS", result)
}

func (s *Service) handleCreateAdminUser(c *fiber.Ctx) error {
	var req CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	err := s.createAdminUser(&req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_CREATE_SUCCESS", nil)
}

func (s *Service) handleUpdateAdminUser(c *fiber.Ctx) error {
	var req UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	userId := c.Params("id")
	err := s.updateAdminUser(userId, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_UPDATE_SUCCESS", nil)
}

func (s *Service) handleDeleteAdminUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	err := s.deleteAdminUser(userId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_DELETE_SUCCESS", nil)
}

func (s *Service) handleToggleAdminUserEnable(c *fiber.Ctx) error {
	userId := c.Params("id")
	err := s.toggleAdminUserEnable(userId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_UPDATE_SUCCESS", nil)
}

func (s *Service) handleGetUsers(c *fiber.Ctx) error {
	result, err := s.getUsers()
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_LIST_SUCCESS", result)
}

func (s *Service) handleGetUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	result, err := s.getUser(userId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_GET_SUCCESS", result)
}

func (s *Service) handleGetAgents(c *fiber.Ctx) error {
	result, err := s.getAgents()
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "AGENT_LIST_SUCCESS", result)
}

func (s *Service) handleGetAgent(c *fiber.Ctx) error {
	agentId := c.Params("id")
	result, err := s.getAgent(agentId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "AGENT_GET_SUCCESS", result)
}

func (s *Service) handleCreateAgent(c *fiber.Ctx) error {
	var req CreateAgentRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	err := s.createAgent(&req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "AGENT_CREATE_SUCCESS", nil)
}

func (s *Service) handleUpdateAgent(c *fiber.Ctx) error {
	var req UpdateAgentRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	agentId := c.Params("id")
	err := s.updateAgent(agentId, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "AGENT_UPDATE_SUCCESS", nil)
}

func (s *Service) handleToggleAgentEnable(c *fiber.Ctx) error {
	agentId := c.Params("id")
	err := s.toggleAgentEnable(agentId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "AGENT_UPDATE_SUCCESS", nil)
}

func (s *Service) handleGetContacts(c *fiber.Ctx) error {
	result, err := s.getContacts()
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "CONTACT_LIST_SUCCESS", result)
}

func (s *Service) handleGetContact(c *fiber.Ctx) error {
	contactId := c.Params("id")
	result, err := s.getContact(contactId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "CONTACT_GET_SUCCESS", result)
}

func (s *Service) handleGetContactAccounts(c *fiber.Ctx) error {
	contactId := c.Params("contactId")
	result, err := s.getContactAccounts(contactId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_LIST_SUCCESS", result)
}

func (s *Service) handleGetDisbursements(c *fiber.Ctx) error {
	var req GetDisbursementFilter

	if err := c.QueryParser(&req); err != nil {
		return err
	}

	result, err := s.getDisbursements(&req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "DISBURSEMENT_LIST_SUCCESS", result)
}

func (s *Service) handleGetDisbursement(c *fiber.Ctx) error {
	disbursementId := c.Params("id")
	result, err := s.getDisbursement(disbursementId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "DISBURSEMENT_GET_SUCCESS", result)
}
