package routers

import (
	"time"

	"angi.id/internal/container"
	m "angi.id/internal/middlewares"
	"angi.id/internal/modules/common"

	// "angi.id/internal/modules/common"
	"angi.id/internal/modules/healthcheck"
	"angi.id/userpb"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Init(app *fiber.App, ctn *container.Container) {
	// app.Get("/health", healthcheck.HealthCheck)
	healthcheck.Init(app, ctn)
	apiV1 := app.Group("/api/v1")
	user := apiV1.Group("/user", m.Auth())
	user.Get("", func(c *fiber.Ctx) error {
		user := &userpb.User{
			Id:    "123",
			Name:  "John Doe",
			Email: "imahihi@gmail.com",
			CreatedAt: &timestamppb.Timestamp{
				Seconds: time.Now().Unix(),
				Nanos:   0,
			},
			UpdatedAt: &timestamppb.Timestamp{
				Seconds: time.Now().Unix(),
				Nanos:   0,
			},
		}
		paging := &userpb.ListUserResponse{
			Users:      []*userpb.User{user, user, user},
			TotalCount: 3,
			Page:       1,
			PageSize:   10,
		}
		res, err := common.EncodeToPositionalJSON(paging)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error encoding user")
		}
		c.Set("Content-Type", "application/json")
		return c.SendString(")]}'\n\n" + string(res))
	})
	// Define your routes here
	// Example:
	// app.Get("/users", func(c *fiber.Ctx) error {
	// 	return c.SendString("List of users")
	// })

	// You can import and use other route files as needed
	// import "angi.id/routers/user"
	// user.SetupUserRoutes(app)
}
