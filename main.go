package main
 
import (
  "fmt"
  "log"

  "github.com/gofiber/fiber/v2"
)

type SiFloor struct {
  ID      int    `json:"id"`
  Height  int    `json:"height"`
  Width   int    `json:"width"`
}

func main() {
  app := fiber.New()

  siFloors := []SiFloor{}

  app.Get("/", func(c *fiber.Ctx) error {
    return c.Status(200).JSON(fiber.Map{"msg": "Welcome to Si-Life!"})
  })

  // Start the server, else log errors
  log.Fatal(app.Listen(":4000"))
  fmt.Println("Server running on port 4000")
}
