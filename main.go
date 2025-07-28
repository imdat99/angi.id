package main

import (
	"angi.account/modules/common"
	protos "angi.account/protos"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func main() {
	app := fiber.New()
	user := &protos.User{
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
	paging := &protos.UserListResponse{
		Users:      []*protos.User{user, user, user},
		TotalCount: 3,
		Page:       1,
		PageSize:   10,
	}
	app.Get("/", func(c *fiber.Ctx) error {
		//return c.SendString("Hello, World!")
		res, err := common.EncodeToPositionalJSON(paging)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error encoding user")
		}
		//c.Set("Content-Type", "application/json+positional-array")
		c.Set("Content-Type", "application/json")
		return c.SendString(")]}'\n\n" + string(res))
	})

	log.Fatal(app.Listen(":3000"))
}
