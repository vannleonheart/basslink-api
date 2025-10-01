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
	var req CreateUserRequest

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
	err := s.updateUser(userId, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_UPDATE_SUCCESS", nil)
}

func (s *Service) handleToggleUserEnable(c *fiber.Ctx) error {
	userId := c.Params("id")
	err := s.toggleUserEnable(userId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_UPDATE_SUCCESS", nil)
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

func (s *Service) handleCreateContact(c *fiber.Ctx) error {
	var req CreateContactRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	err := s.createContact(agentUser.Agent, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "CONTACT_CREATE_SUCCESS", nil)
}

func (s *Service) handleUpdateContact(c *fiber.Ctx) error {
	var req UpdateContactRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	contactId := c.Params("id")
	err := s.updateContact(contactId, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "CONTACT_UPDATE_SUCCESS", nil)
}

func (s *Service) handleDeleteContact(c *fiber.Ctx) error {
	contactId := c.Params("id")
	err := s.deleteContact(contactId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "CONTACT_DELETE_SUCCESS", nil)
}

func (s *Service) handleCreateContactDocument(c *fiber.Ctx) error {
	var req CreateContactDocumentRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	contactId := c.Params("contactId")
	err := s.createContactDocument(contactId, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "CONTACT_DOCUMENT_CREATE_SUCCESS", nil)
}

func (s *Service) handleUpdateContactDocument(c *fiber.Ctx) error {
	var req UpdateContactDocumentRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	contactId := c.Params("contactId")
	documentId := c.Params("documentId")
	err := s.updateContactDocument(contactId, documentId, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "CONTACT_DOCUMENT_UPDATE_SUCCESS", nil)
}

func (s *Service) handleDeleteContactDocument(c *fiber.Ctx) error {
	contactId := c.Params("contactId")
	documentId := c.Params("documentId")
	err := s.deleteContactDocument(contactId, documentId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "CONTACT_DOCUMENT_DELETE_SUCCESS", nil)
}

func (s *Service) handleGetContactAccounts(c *fiber.Ctx) error {
	contactId := c.Params("contactId")
	result, err := s.getContactAccounts(contactId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "USER_LIST_SUCCESS", result)
}

func (s *Service) handleCreateContactAccount(c *fiber.Ctx) error {
	var req CreateContactAccountRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	contactId := c.Params("contactId")
	err := s.createContactAccount(contactId, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "CONTACT_ACCOUNT_CREATE_SUCCESS", nil)
}

func (s *Service) handleUpdateContactAccount(c *fiber.Ctx) error {
	var req UpdateContactAccountRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	contactId := c.Params("contactId")
	accountId := c.Params("accountId")
	err := s.updateContactAccount(contactId, accountId, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "CONTACT_ACCOUNT_UPDATE_SUCCESS", nil)
}

func (s *Service) handleDeleteContactAccount(c *fiber.Ctx) error {
	contactId := c.Params("contactId")
	accountId := c.Params("accountId")
	err := s.deleteContactAccount(contactId, accountId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "CONTACT_ACCOUNT_DELETE_SUCCESS", nil)
}

func (s *Service) handleGetDisbursements(c *fiber.Ctx) error {
	var req GetDisbursementFilter

	if err := c.QueryParser(&req); err != nil {
		return err
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	result, err := s.getDisbursements(agentUser.Agent, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "DISBURSEMENT_LIST_SUCCESS", result)
}

func (s *Service) handleGetDisbursement(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	disbursementId := c.Params("id")
	result, err := s.getDisbursement(agentUser.Agent, disbursementId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "DISBURSEMENT_GET_SUCCESS", result)
}

func (s *Service) handleCreateDisbursement(c *fiber.Ctx) error {
	var req CreateDisbursementRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	err := s.createDisbursement(agentUser.Agent, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "DISBURSEMENT_CREATE_SUCCESS", nil)
}
