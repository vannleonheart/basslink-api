package admin

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

	adminUser := c.Locals("admin").(*basslink.Administrator)
	err := s.updatePassword(adminUser, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleGetUsers(c *fiber.Ctx) error {
	result, err := s.getUsers()
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleGetUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	result, err := s.getUser(userId)
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

	err := s.createUser(&req)
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

	userId := c.Params("id")
	err := s.updateUser(userId, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleDeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	err := s.deleteUser(userId)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleToggleUserEnable(c *fiber.Ctx) error {
	userId := c.Params("id")
	err := s.toggleUserEnable(userId)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleGetAgents(c *fiber.Ctx) error {
	result, err := s.getAgents()
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleGetAgent(c *fiber.Ctx) error {
	agentId := c.Params("id")
	result, err := s.getAgent(agentId)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleCreateAgent(c *fiber.Ctx) error {
	var req CreateAgentRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	err := s.createAgent(&req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (s *Service) handleUpdateAgent(c *fiber.Ctx) error {
	var req UpdateAgentRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		return err
	}

	agentId := c.Params("id")
	err := s.updateAgent(agentId, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleToggleAgentEnable(c *fiber.Ctx) error {
	agentId := c.Params("id")
	err := s.toggleAgentEnable(agentId)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleGetDisbursements(c *fiber.Ctx) error {
	result, err := s.getDisbursements()
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleGetDisbursement(c *fiber.Ctx) error {
	disbursementId := c.Params("id")
	result, err := s.getDisbursement(disbursementId)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
