package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kudaibergenoff/fiber-gorm-app/app/handlers"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	user := api.Group("/users")
	// user routes
	user.Get("/", handlers.GetAllUsers)
}
