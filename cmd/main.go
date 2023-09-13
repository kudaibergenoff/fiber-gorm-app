package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kudaibergenoff/fiber-gorm-app/database"
	"github.com/kudaibergenoff/fiber-gorm-app/router"
	_ "github.com/lib/pq"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetupRoutes(app)
	// handle unavailable route
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
