package server

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)

	s.App.Get("/health", s.healthHandler)

	s.App.Get("/auth/:provider/callback", s.getAuthCallbackFunction)
	s.App.Get("/logout/:provider", s.getAuthLogoutFunction)
	s.App.Get("/auth/:provider", goth_fiber.BeginAuthHandler)
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

func (s *FiberServer) getAuthCallbackFunction(c *fiber.Ctx) error {
	provider := c.Params("provider")
	context.Background().Value(provider)

	user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		fmt.Fprintln(c, err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	fmt.Println(user)
	return c.Redirect("http://127.0.0.1:5173", fiber.StatusTemporaryRedirect)
}

func (s *FiberServer) getAuthLogoutFunction(c *fiber.Ctx) error {
	provider := c.Params("provider")
	context.Background().Value(provider)

	goth_fiber.Logout(c)
	return c.Redirect("http://127.0.0.1:5173", fiber.StatusTemporaryRedirect)
}

// func (s *FiberServer) getAuthHandler(c *fiber.Ctx) error {
// 	provider := c.Params("provider")
// 	context.Background().Value(provider)
//
// 	// Try to get the user without re-authenticating
// 	if user, err := goth_fiber.CompleteUserAuth(c); err == nil {
// 		c.JSON(user)
// 	} else {
// 		goth_fiber.BeginAuthHandler(c)
// 	}
// }
