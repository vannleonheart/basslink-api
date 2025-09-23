package user

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

	user := c.Locals("user").(*basslink.User)
	err := s.updatePassword(user, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleGetContacts(c *fiber.Ctx) error {
	user := c.Locals("user").(*basslink.User)
	result, err := s.getContacts(user)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleGetContact(c *fiber.Ctx) error {
	user := c.Locals("user").(*basslink.User)
	contactId := c.Params("id")
	result, err := s.getContact(user, contactId)
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

	user := c.Locals("user").(*basslink.User)
	err := s.createContact(user, &req)
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

	user := c.Locals("user").(*basslink.User)
	contactId := c.Params("id")
	err := s.updateContact(user, contactId, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleDeleteContact(c *fiber.Ctx) error {
	user := c.Locals("user").(*basslink.User)
	contactId := c.Params("id")
	err := s.deleteContact(user, contactId)
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
	user := c.Locals("user").(*basslink.User)
	err := s.createContactDocument(user, contactId, &req)
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

	user := c.Locals("user").(*basslink.User)
	contactId := c.Params("contactId")
	documentId := c.Params("documentId")
	err := s.updateContactDocument(user, contactId, documentId, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleDeleteContactDocument(c *fiber.Ctx) error {
	user := c.Locals("user").(*basslink.User)
	contactId := c.Params("contactId")
	documentId := c.Params("documentId")
	err := s.deleteContactDocument(user, contactId, documentId)
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
	user := c.Locals("user").(*basslink.User)
	err := s.createContactAccount(user, contactId, &req)
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

	user := c.Locals("user").(*basslink.User)
	contactId := c.Params("contactId")
	accountId := c.Params("accountId")
	err := s.updateContactAccount(user, contactId, accountId, &req)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleDeleteContactAccount(c *fiber.Ctx) error {
	user := c.Locals("user").(*basslink.User)
	contactId := c.Params("contactId")
	accountId := c.Params("accountId")
	err := s.deleteContactAccount(user, contactId, accountId)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (s *Service) handleGetDisbursements(c *fiber.Ctx) error {
	user := c.Locals("user").(*basslink.User)
	result, err := s.getDisbursements(user)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Service) handleGetDisbursement(c *fiber.Ctx) error {
	user := c.Locals("user").(*basslink.User)
	disbursementId := c.Params("id")
	result, err := s.getDisbursement(user, disbursementId)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
