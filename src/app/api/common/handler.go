package common

import (
	"CRM/src/lib/basslink"
	"encoding/json"
	"errors"

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

	result, apErr := s.handleUploadFile(path, form)
	if apErr != nil {
		return apErr
	}

	return basslink.NewSuccessResponse(c, "FILE_UPLOAD_SUCCESS", result)
}

func (s *Service) PublicFileUploadHandler(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	path := "uploads/general"

	result, apErr := s.handleUploadFile(path, form)
	if apErr != nil {
		return apErr
	}

	return basslink.NewSuccessResponse(c, "FILE_UPLOAD_SUCCESS", result)
}

func (s *Service) GetCurrenciesHandler(c *fiber.Ctx) error {
	result, err := s.handleGetCurrencies()
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "CURRENCY_LIST_SUCCESS", result)
}

func (s *Service) GetRateHandler(c *fiber.Ctx) error {
	var req GetRateRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	result, apErr := s.handleGetRate(&req)
	if apErr != nil {
		return apErr
	}

	return basslink.NewSuccessResponse(c, "RATES_GET_SUCCESS", result)
}

func (s *Service) CreateAppointmentHandler(c *fiber.Ctx) error {
	var req CreateAppointmentRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	ipAddress := c.IP()
	isRecaptchaValid, err := s.App.Recaptcha.Verify(req.Token, &ipAddress)
	if err != nil {
		return err
	}

	if !isRecaptchaValid {
		return errors.New("invalid recaptcha")
	}

	result, err := s.handleCreateAppointment(&req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "APPOINTMENT_CREATE_SUCCESS", result)
}

func (s *Service) SearchTransactionHandler(c *fiber.Ctx) error {
	var req TransactionSearchRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	result, err := s.searchTransaction(&req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "TRANSACTION_SEARCH_SUCCESS", result)
}

func (s *Service) CreateTransactionHandler(c *fiber.Ctx) error {
	var req CreateRemittanceRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	ipAddress := c.IP()
	isRecaptchaValid, err := s.App.Recaptcha.Verify(req.Token, &ipAddress)
	if err != nil {
		return err
	}

	if !isRecaptchaValid {
		return errors.New("invalid recaptcha")
	}

	result, err := s.createTransaction(&req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "TRANSACTION_CREATE_SUCCESS", result)
}

func (s *Service) GetPaymentHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	result, apErr := s.handleGetPaymentById(id)
	if apErr != nil {
		return apErr
	}

	return basslink.NewSuccessResponse(c, "PAYMENT_GET_SUCCESS", result)
}

func (s *Service) ConfirmPaymentHandler(c *fiber.Ctx) error {
	var req PaymentConfirmRequest

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := s.App.ValidateRequest(&req); err != nil {
		var errorData []map[string]interface{}

		_ = json.Unmarshal([]byte(err.Error()), &errorData)

		return basslink.NewAppError("ERROR_VALIDATION", basslink.ErrBadRequest, basslink.ErrBadRequestValidation, "", errorData)
	}

	ipAddress := c.IP()
	isRecaptchaValid, err := s.App.Recaptcha.Verify(req.Token, &ipAddress)
	if err != nil {
		return err
	}

	if !isRecaptchaValid {
		return errors.New("invalid recaptcha")
	}

	id := c.Params("id")
	err = s.handlePaymentConfirm(id, &req)
	if err != nil {
		return err
	}

	return basslink.NewSuccessResponse(c, "PAYMENT_CONFIRM_SUCCESS", nil)
}
