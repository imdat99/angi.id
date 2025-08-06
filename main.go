package main

import (
	"fmt"
	"log"

	"angi.id/internal/container"
	"angi.id/internal/routers"
	"angi.id/internal/shared"
	"angi.id/internal/shared/db"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(shared.FiberConfig())
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	dbpool, err := db.ConnectPostgres()
	redisClient := db.NewRedisClient()
	container := container.NewContainer(dbpool, redisClient)

	if err != nil {
		log.Fatalf("Error connecting to Postgres: %v", err)
		return
	}

	defer redisClient.Close()
	defer dbpool.Close()

	routers.Init(app, container)

	// notfoundary
	app.Use(shared.NotFoundHandler)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", shared.Acfg.Server.Port)))
}
