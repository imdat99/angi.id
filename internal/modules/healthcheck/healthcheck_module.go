package healthcheck

import (
	c "angi.id/internal/container"
	"github.com/gofiber/fiber/v2"
)

func Init(app fiber.Router, ctn *c.Container) {
	controller := newHealthCheckController(HewHealthCheckService(ctn))
	app.Get("/health", controller.Check)
}
