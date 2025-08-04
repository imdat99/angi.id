package response

import "github.com/gofiber/fiber/v2"

type Common struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessWithPaginate[T any] struct {
	Code         int    `json:"code"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	Results      []T    `json:"results"`
	Page         int    `json:"page"`
	Limit        int    `json:"limit"`
	TotalPages   int64  `json:"total_pages"`
	TotalResults int64  `json:"total_results"`
}

type ErrorDetails struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}
type responseWithData[T any] struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

const (
	SuccessCode   = 200
	ErrorCode     = 500
	SuccessStatus = "success"
	ErrorStatus   = "error"
)

func Ok[T any](c *fiber.Ctx, message string, data T) error {
	return c.Status(SuccessCode).JSON(responseWithData[T]{
		Code:    SuccessCode,
		Status:  SuccessStatus,
		Message: message,
		Data:    data,
	})
}
