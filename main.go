package main

import (
	"fmt"
	"log"

	"angi.id/internal/response"
	"angi.id/internal/routers"
	"angi.id/internal/shared/config"
	"angi.id/internal/shared/db"
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

	dbpool, err := db.ConnectPostgres()
	redisClient := db.NewRedisClient()

	if err != nil {
		log.Fatalf("Error connecting to Postgres: %v", err)
		return
	}

	defer redisClient.Close()
	defer dbpool.Close()

	routers.Init(app)

	// notfoundary
	app.Use(func(c *fiber.Ctx) error {
		return response.Error(c, fiber.StatusNotFound, "Endpoint Not Found", nil)
	})
	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Acfg.Server.Port)))
}
