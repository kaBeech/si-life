package main
 
import (
  "fmt"
  "log"

  "github.com/gofiber/fiber/v2"
)
 
func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.Status(200).JSON(fiber.Map{"msg": "Welcome to Si-Life!"})
  })

  // Start the server, else log errors
  log.Fatal(app.Listen(":4000"))
  fmt.Println("Server running on port 4000")
}
