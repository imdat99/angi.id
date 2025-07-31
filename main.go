package main

import (
	"log"
	"time"

	"angi.account/response"
	"angi.account/routers"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"angi.account/modules/common"
	"angi.account/userpb"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	routers.SetupRoutes(app)

	// notfoundary
	app.Use(func(c *fiber.Ctx) error {
		return response.Error(c, fiber.StatusNotFound, "Endpoint Not Found", nil)
	})

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
	app.Get("/", func(c *fiber.Ctx) error {
		res, err := common.EncodeToPositionalJSON(paging)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error encoding user")
		}
		c.Set("Content-Type", "application/json")
		return c.SendString(")]}'\n\n" + string(res))
	})
	log.Fatal(app.Listen(":3000"))
}
