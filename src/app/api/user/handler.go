package user

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
	user := c.Locals("user").(*basslink.User)
	result, err := s.getProfile(user)
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

	user := c.Locals("user").(*basslink.User)
	err := s.updatePassword(user, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "ACCOUNT_PASSWORD_UPDATE_SUCCESS", nil)
}

func (s *Service) handleGetDisbursements(c *fiber.Ctx) error {
	user := c.Locals("user").(*basslink.User)
	result, err := s.getDisbursements(user)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "DISBURSEMENT_LIST_SUCCESS", result)
}

func (s *Service) handleGetDisbursement(c *fiber.Ctx) error {
	user := c.Locals("user").(*basslink.User)
	disbursementId := c.Params("id")
	result, err := s.getDisbursement(user, disbursementId)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "DISBURSEMENT_GET_SUCCESS", result)
}
