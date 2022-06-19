package middlewares

import (
	"go-api/configs"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	if authorization == configs.ApiKey() {
		return c.Next()
	}

	return c.Status(http.StatusBadRequest).JSON("Authorization required")
}
