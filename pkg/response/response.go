package response

import "github.com/gofiber/fiber/v3"

func Success(c fiber.Ctx, data any) error {
	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func Error(c fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}