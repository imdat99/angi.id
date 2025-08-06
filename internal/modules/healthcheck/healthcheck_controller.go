package healthcheck

import (
	"angi.id/internal/response"
	"github.com/gofiber/fiber/v2"
)

type healthCheckController struct {
	HealthCheckService HealthCheckService
}
type HealthCheckController interface {
	Check(c *fiber.Ctx) error
}

func newHealthCheckController(healthCheckService HealthCheckService) HealthCheckController {
	return &healthCheckController{
		HealthCheckService: healthCheckService,
	}
}
func (h *healthCheckController) addServiceStatus(
	serviceList *[]response.HealthCheck, name string, isUp bool, message *string,
) {
	status := "Up"

	if !isUp {
		status = "Down"
	}

	*serviceList = append(*serviceList, response.HealthCheck{
		Name:    name,
		Status:  status,
		IsUp:    isUp,
		Message: message,
	})
}
func (h *healthCheckController) Check(c *fiber.Ctx) error {
	isHealthy := true
	var serviceList []response.HealthCheck

	// Check the database connection
	if err := h.HealthCheckService.DBcheck(); err != nil {
		isHealthy = false
		errMsg := err.Error()
		h.addServiceStatus(&serviceList, "Postgre", false, &errMsg)
	} else {
		h.addServiceStatus(&serviceList, "Postgre", true, nil)
	}

	if err := h.HealthCheckService.MemoryHeapCheck(); err != nil {
		isHealthy = false
		errMsg := err.Error()
		h.addServiceStatus(&serviceList, "Memory", false, &errMsg)
	} else {
		h.addServiceStatus(&serviceList, "Memory", true, nil)
	}
	if err := h.HealthCheckService.RedisCheck(); err != nil {
		isHealthy = false
		errMsg := err.Error()
		h.addServiceStatus(&serviceList, "Redis", false, &errMsg)
	} else {
		h.addServiceStatus(&serviceList, "Redis", true, nil)
	}
	// Return the response based on health check result
	statusCode := fiber.StatusOK
	status := "success"

	if !isHealthy {
		statusCode = fiber.StatusInternalServerError
		status = "error"
	}

	return c.Status(statusCode).JSON(response.HealthCheckResponse{
		Status:    status,
		Message:   "Health check completed",
		Code:      statusCode,
		IsHealthy: isHealthy,
		Result:    serviceList,
	})
}
