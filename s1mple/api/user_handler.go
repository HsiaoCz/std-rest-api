package api

import (
	"github.com/HsiaoCz/std-rest-api/s1mple/types"
	"github.com/gofiber/fiber/v2"
)

// HandleCreateUser
// handle /api/v1/user
func HandleCreateUser(c *fiber.Ctx) error {
	post := new(types.PostUser)
	if err := c.BodyParser(post); err != nil {
		return err
	}
	return nil
}

// HandleGetUserById
// handle /api/v1/user/:id
func HandleGetUserById(c *fiber.Ctx) error {
	return nil
}
