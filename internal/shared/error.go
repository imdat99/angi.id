package shared

import (
	"errors"

	"angi.id/internal/response"
	"angi.id/pkg/validation"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if errorsMap := validation.CustomErrorMessages(err); len(errorsMap) > 0 {
		return response.Error(c, fiber.StatusBadRequest, "Bad Request", errorsMap)
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return response.Error(c, fiberErr.Code, fiberErr.Message, nil)
	}

	return response.Error(c, fiber.StatusInternalServerError, "Internal Server Error", nil)
}

func NotFoundHandler(c *fiber.Ctx) error {
	return response.Error(c, fiber.StatusNotFound, "Endpoint Not Found", nil)
}
