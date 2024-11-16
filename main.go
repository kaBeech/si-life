package main
 
import (
  "fmt"
  "log"
  "os"

  "github.com/gofiber/fiber/v2"
  "github.com/joho/godotenv"
)

type SiFloor struct {
  ID      int    `json:"id"`
  Height  int    `json:"height"`
  Width   int    `json:"width"`
}

func main() {
  app := fiber.New()

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  PORT := os.Getenv("PORT")

  siFloors := []SiFloor{}

  // Index route
  app.Get("/", func(c *fiber.Ctx) error {
    return c.Status(200).JSON(fiber.Map{"msg": "Welcome to  SiLife!"})
  })

  // Get all Si-Floors
  app.Get("/api/sifloor", func(c *fiber.Ctx) error {
    return c.JSON(siFloors)
  })

  // Create a new Si-Floor
  app.Post("/api/sifloor", func(c *fiber.Ctx) error {
    siFloor := &SiFloor{}
  

    if err := c.BodyParser(siFloor); err != nil {
      return err
    }

    if siFloor.Height == 0 || siFloor.Width == 0 {
      return c.Status(400).JSON(fiber.Map{"error": "Height and Width are required"})
    }

    siFloor.ID = len(siFloors) + 1
    siFloors = append(siFloors, *siFloor)

    return c.Status(201).JSON(siFloor)
  })

  // Start the server, else log errors
  log.Fatal(app.Listen(":" + PORT))
  fmt.Println("Server running on port 4000")
}
