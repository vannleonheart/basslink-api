package basslink

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string      `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewErrorResponse(c *fiber.Ctx, err *AppError) error {
	statusCode := fiber.StatusInternalServerError
	resp := Response{
		Status:  "error",
		Code:    fmt.Sprintf("%s.%s", err.Code, err.Kind),
		Message: err.Error(),
		Data:    err.Data,
	}

	code, e := strconv.Atoi(err.Code)
	if e == nil {
		statusCode = code
	}

	return c.Status(statusCode).JSON(resp)
}

func NewSuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	statusCode := fiber.StatusOK
	resp := Response{
		Status:  "success",
		Code:    "200.00",
		Message: message,
		Data:    data,
	}

	return c.Status(statusCode).JSON(resp)
}
