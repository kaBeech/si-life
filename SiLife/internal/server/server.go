package server

import (
	"github.com/gofiber/fiber/v2"

	"SiLife/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "SiLife",
			AppName:      "SiLife",
		}),

		db: database.New(),
	}

	return server
}
