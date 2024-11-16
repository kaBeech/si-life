package main
 
import (
  // "context"
  "fmt"
  "log"
  "os"

  "github.com/gofiber/fiber/v2"
  "github.com/joho/godotenv"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

type SiFloor struct {
  gorm.Model
  ID      int    `json:"id"`
  Height  int    `json:"height"`
  Width   int    `json:"width"`
}

var db *gorm.DB

func main() {
  errGoDotEnv := godotenv.Load()
  if errGoDotEnv != nil {
    log.Fatal("Error loading .env file")
  }

  PORT := os.Getenv("PORT")
  DB_URL := os.Getenv("DB_URL")

  dsn := DB_URL
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&SiFloor{})

  fmt.Println("Welcome to SiLife!")

  app := fiber.New()

  // Index route
  app.Get("/", getHome)

  // Create a new Si-Floor
  app.Post("/api/sifloor", func(c *fiber.Ctx) error {
    siFloor := &SiFloor{}


    if err := c.BodyParser(siFloor); err != nil {
      return err
    }

    if siFloor.Height < 1 || siFloor.Width < 1 {
      return c.Status(400).JSON(fiber.Map{"error": "Height and Width must be greater than 0"})
    }
    db.Create(&SiFloor{Height: siFloor.Height, Width: siFloor.Width})
    return c.Status(201).JSON(fiber.Map{"msg": "SiFloor created successfully"})
  })

  // Count SiFloors
  app.Get("/api/sifloor", func(c *fiber.Ctx) error {
    result := db.Find(&siFloor)

    if result.Error != nil {
      panic("failed to list SiFloors")
    }
    message := "There are " + fmt.Sprint(result.RowsAffected) + " SiFloors currently in the database"
    return c.Status(200).JSON(fiber.Map{"msg": message})
  })

  // Get a SiFloor by ID
  app.Get("/api/sifloor/:id", func(c *fiber.Ctx) error {
    id := c.Params("id")
    var siFloor SiFloor
    db.First(&siFloor, id)
    // db.First(&siFloor, "height = ?", 10) // find SiFloor with height 10
    return c.Status(200).JSON(siFloor)
  })

  // Update a SiFloor
  app.Put("/api/sifloor/:id", func(c *fiber.Ctx) error {
    id := c.Params("id")
    var siFloor SiFloor
    db.First(&siFloor, id)
    db.Model(&siFloor).Updates(SiFloor{Height: 20, Width: 20}) // non-zero fields
    return c.Status(200).JSON(fiber.Map{"msg": "SiFloor updated successfully"})
  })

  // Delete a SiFloor
  app.Delete("/api/sifloor/:id", func(c *fiber.Ctx) error {
    id := c.Params("id")
    var siFloor SiFloor
    db.First(&siFloor, id)
    db.Delete(&siFloor)
    return c.Status(200).JSON(fiber.Map{"msg": "SiFloor deleted successfully"})
  })

  // Start the server, else log errors
  log.Fatal(app.Listen(":" + PORT))
  fmt.Println("Server running on port 4000")
}

func getHome(c *fiber.Ctx) error {
  return c.Status(200).JSON(fiber.Map{"msg": "Welcome to  SiLife!"})
}
