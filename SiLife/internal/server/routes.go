package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)

	s.App.Get("/health", s.healthHandler)

	s.App.Get("/auth/:provider/callback", s.getAuthCallbackFunction)
	s.App.Get("/logout/:provider", s.getAuthLogoutFunction)
	s.App.Get("/auth/:provider", s.getAuthHandler)
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
	provider, err := c.Params("provider")
	c.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(c)
	if err != nil {
		fmt.Fprintln(c, err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	fmt.Println(user)
	http.Redirect(c, c.Request(), "http://localhost:5173", http.StatusFound)
}

func (s *FiberServer) getAuthLogoutFunction(c *fiber.Ctx) error {
	provider, err := c.Params("provider")
	c.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.Logout(c)
	http.Redirect(c, c.Request(), "http://localhost:5173", http.StatusFound)
}

func (s *FiberServer) getAuthHandler(c *fiber.Ctx) error {
	provider, err := c.Params("provider")
	c.WithContext(context.WithValue(context.Background(), "provider", provider))

	// Try to get the user without re-authenticating
	if user, err := gothic.CompleteUserAuth(c); err == nil {
		c.JSON(user)
	} else {
		gothic.BeginAuthHandler(c)
	}
}
