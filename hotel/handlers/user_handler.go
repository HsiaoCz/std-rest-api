package handlers

import (
	"net/http"

	"github.com/HsiaoCz/std-rest-api/hotel/types"
	"github.com/gofiber/fiber/v2"
)

// handle api/v1/user
// HandleGetUser
func HandleGetUser(c *fiber.Ctx) error {
	user := &types.User{}
	if err := c.BodyParser(user); err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

// handle api/v1/user:id
// HandleGetUserByID
func HandleGetUserByID(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"msg": "james",
	})
}
