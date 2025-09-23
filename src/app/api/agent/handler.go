package agent

import (
	"CRM/src/lib/basslink"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) handleSignIn(c *fiber.Ctx) error {
	var req SignInRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	result, err := s.signIn(&req)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleUpdatePassword(c *fiber.Ctx) error {
	var req UpdatePasswordRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	err := s.updatePassword(agentUser, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleGetAgentUsers(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	result, err := s.getAgentUsers(agentUser.Agent)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleGetAgentUser(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	agentUserId := c.Params("id")
	result, err := s.getAgentUser(agentUser.Agent, agentUserId)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleCreateAgentUser(c *fiber.Ctx) error {
	var req CreateAgentUserRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	err := s.createAgentUser(agentUser.Agent, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (s *Service) handleUpdateAgentUser(c *fiber.Ctx) error {
	var req UpdateAgentUserRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	agentUserId := c.Params("id")
	err := s.updateAgentUser(agentUser.Agent, agentUserId, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleDeleteAgentUser(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	agentUserId := c.Params("id")
	err := s.deleteAgentUser(agentUser.Agent, agentUserId)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleToggleAgentUserEnable(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	agentUserId := c.Params("id")
	err := s.toggleAgentUserEnable(agentUser.Agent, agentUserId)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleGetUsers(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	result, err := s.getUsers(agentUser.Agent)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleGetUser(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	userId := c.Params("id")
	result, err := s.getAgentUser(agentUser.Agent, userId)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleCreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	err := s.createUser(agentUser.Agent, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (s *Service) handleUpdateUser(c *fiber.Ctx) error {
	var req UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	userId := c.Params("id")
	err := s.updateUser(agentUser.Agent, userId, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleToggleUserEnable(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	userId := c.Params("id")
	err := s.toggleUserEnable(agentUser.Agent, userId)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleGetContacts(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	result, err := s.getContacts(agentUser.Agent)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleGetContact(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	contactId := c.Params("id")
	result, err := s.getContact(agentUser.Agent, contactId)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleCreateContact(c *fiber.Ctx) error {
	var req CreateContactRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	err := s.createContact(agentUser.Agent, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (s *Service) handleUpdateContact(c *fiber.Ctx) error {
	var req UpdateContactRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	contactId := c.Params("id")
	err := s.updateContact(agentUser.Agent, contactId, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleDeleteContact(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	contactId := c.Params("id")
	err := s.deleteContact(agentUser.Agent, contactId)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleCreateContactDocument(c *fiber.Ctx) error {
	var req CreateContactDocumentRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	contactId := c.Params("contactId")
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	err := s.createContactDocument(agentUser.Agent, contactId, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (s *Service) handleUpdateContactDocument(c *fiber.Ctx) error {
	var req UpdateContactDocumentRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	contactId := c.Params("contactId")
	documentId := c.Params("documentId")
	err := s.updateContactDocument(agentUser.Agent, contactId, documentId, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleDeleteContactDocument(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	contactId := c.Params("contactId")
	documentId := c.Params("documentId")
	err := s.deleteContactDocument(agentUser.Agent, contactId, documentId)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleCreateContactAccount(c *fiber.Ctx) error {
	var req CreateContactAccountRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	contactId := c.Params("contactId")
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	err := s.createContactAccount(agentUser.Agent, contactId, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (s *Service) handleUpdateContactAccount(c *fiber.Ctx) error {
	var req UpdateContactAccountRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	agentUser := c.Locals("agent").(*basslink.AgentUser)
	contactId := c.Params("contactId")
	accountId := c.Params("accountId")
	err := s.updateContactAccount(agentUser.Agent, contactId, accountId, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleDeleteContactAccount(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	contactId := c.Params("contactId")
	accountId := c.Params("accountId")
	err := s.deleteContactAccount(agentUser.Agent, contactId, accountId)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleGetDisbursements(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	result, err := s.getDisbursements(agentUser.Agent)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleGetDisbursement(c *fiber.Ctx) error {
	agentUser := c.Locals("agent").(*basslink.AgentUser)
	disbursementId := c.Params("id")
	result, err := s.getDisbursement(agentUser.Agent, disbursementId)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
