package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kudaibergenoff/fiber-gorm-app/app/models"
	"github.com/kudaibergenoff/fiber-gorm-app/database"
)

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []models.User
	db.Find(&users)
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
	// return users
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users Found", "data": users})
}
