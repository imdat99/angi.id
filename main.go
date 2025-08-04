package main

import (
	"fmt"
	"log"

	"angi.account/config"
	"angi.account/modules/storage"
	"angi.account/response"
	"angi.account/routers"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	dbpool, err := storage.ConnectMySQL()
	if err != nil {
		log.Fatalf("Error connecting to MySQL: %v", err)
		return
	}
	defer dbpool.Close()

	routers.SetupRoutes(app)

	// notfoundary
	app.Use(func(c *fiber.Ctx) error {
		return response.Error(c, fiber.StatusNotFound, "Endpoint Not Found", nil)
	})
	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Acfg.Server.Port)))
}
