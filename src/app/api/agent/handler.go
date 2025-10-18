package agent

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
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	result, err := s.getProfile(agentUser)
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

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	err := s.updatePassword(agentUser, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "ACCOUNT_PASSWORD_UPDATE_SUCCESS", nil)
}

func (s *Service) handleGetAgentUsers(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	result, err := s.getAgentUsers(agentUser.Agent)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_LIST_SUCCESS", result)
}

func (s *Service) handleGetAgentUser(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	agentUserId := c.Params("id")
	result, err := s.getAgentUser(agentUser.Agent, agentUserId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_GET_SUCCESS", result)
}

func (s *Service) handleCreateAgentUser(c *fiber.Ctx) error {
	var req CreateAgentUserRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	err := s.createAgentUser(agentUser.Agent, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_CREATE_SUCCESS", nil)
}

func (s *Service) handleUpdateAgentUser(c *fiber.Ctx) error {
	var req UpdateAgentUserRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	agentUserId := c.Params("id")
	err := s.updateAgentUser(agentUser.Agent, agentUserId, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_UPDATE_SUCCESS", nil)
}

func (s *Service) handleDeleteAgentUser(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	agentUserId := c.Params("id")
	err := s.deleteAgentUser(agentUser.Agent, agentUserId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_DELETE_SUCCESS", nil)
}

func (s *Service) handleToggleAgentUserEnable(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	agentUserId := c.Params("id")
	err := s.toggleAgentUserEnable(agentUser.Agent, agentUserId)
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

func (s *Service) handleCreateUser(c *fiber.Ctx) error {
	var req CreateSenderRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	err := s.createUser(agentUser.Agent, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_CREATE_SUCCESS", nil)
}

func (s *Service) handleUpdateUser(c *fiber.Ctx) error {
	var req UpdateSenderRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	userId := c.Params("id")
	err := s.updateUser(agentUser, userId, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_UPDATE_SUCCESS", nil)
}

func (s *Service) handleGetRecipients(c *fiber.Ctx) error {
	result, err := s.getRecipients()
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "RECIPIENT_LIST_SUCCESS", result)
}

func (s *Service) handleGetRecipient(c *fiber.Ctx) error {
	recipientId := c.Params("id")
	result, err := s.getRecipient(recipientId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "RECIPIENT_GET_SUCCESS", result)
}

func (s *Service) handleCreateRecipient(c *fiber.Ctx) error {
	var req CreateRecipientRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	err := s.createRecipient(agentUser.Agent, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "RECIPIENT_CREATE_SUCCESS", nil)
}

func (s *Service) handleUpdateRecipient(c *fiber.Ctx) error {
	var req UpdateRecipientRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	recipientId := c.Params("id")
	err := s.updateRecipient(recipientId, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "RECIPIENT_UPDATE_SUCCESS", nil)
}

func (s *Service) handleDeleteRecipient(c *fiber.Ctx) error {
	recipientId := c.Params("id")
	err := s.deleteRecipient(recipientId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "RECIPIENT_DELETE_SUCCESS", nil)
}

func (s *Service) handleCreateRecipientDocument(c *fiber.Ctx) error {
	var req CreateRecipientDocumentRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	recipientId := c.Params("recipientId")
	err := s.createRecipientDocument(recipientId, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "RECIPIENT_DOCUMENT_CREATE_SUCCESS", nil)
}

func (s *Service) handleUpdateRecipientDocument(c *fiber.Ctx) error {
	var req UpdateRecipientDocumentRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	recipientId := c.Params("recipientId")
	documentId := c.Params("documentId")
	err := s.updateRecipientDocument(recipientId, documentId, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "RECIPIENT_DOCUMENT_UPDATE_SUCCESS", nil)
}

func (s *Service) handleDeleteRecipientDocument(c *fiber.Ctx) error {
	recipientId := c.Params("recipientId")
	documentId := c.Params("documentId")
	err := s.deleteRecipientDocument(recipientId, documentId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "RECIPIENT_DOCUMENT_DELETE_SUCCESS", nil)
}

func (s *Service) handleGetRemittances(c *fiber.Ctx) error {
	var req GetRemittanceFilter

	if err := c.QueryParser(&req); err != nil {
		return err
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	result, err := s.getRemittances(agentUser.Agent, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "REMITTANCE_LIST_SUCCESS", result)
}

func (s *Service) handleGetRemittance(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	remittanceId := c.Params("id")
	result, err := s.getRemittance(agentUser.Agent, remittanceId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "REMITTANCE_GET_SUCCESS", result)
}

func (s *Service) handleCreateRemittance(c *fiber.Ctx) error {
	var req CreateRemittanceRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	err := s.createRemittance(agentUser.Agent, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "REMITTANCE_CREATE_SUCCESS", nil)
}
