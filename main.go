package main
 
import (
  "context"
  "fmt"
  "log"
  "os"

  "github.com/gofiber/fiber/v2"
  "github.com/jackc/pgx/v5"
  "github.com/joho/godotenv"
)

type SiFloor struct {
  ID      int    `json:"id"`
  Height  int    `json:"height"`
  Width   int    `json:"width"`
}

func main() {
  errGoDotEnv := godotenv.Load()
  if errGoDotEnv != nil {
    log.Fatal("Error loading .env file")
  }

  conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
  if err != nil {
  	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
  	os.Exit(1)
  }
  defer conn.Close(context.Background())

  var greeting string
  err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
  if err != nil {
  	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
  	os.Exit(1)
  }

  fmt.Println(greeting)

  app := fiber.New()

  PORT := os.Getenv("PORT")

  siFloors := []SiFloor{}

  app.Get("/", getHome)

  app.Get("/api/sifloor", getSiFloors)

  app.Post("/api/sifloor", createSiFloor)

  // Index route
  func getHome(c *fiber.Ctx) error {
    return c.Status(200).JSON(fiber.Map{"msg": "Welcome to  SiLife!"})
  }

  // Get all Si-Floors
  func getSiFloors(c *fiber.Ctx) error {
    return c.JSON(siFloors)
  }

  // Create a new Si-Floor
  func createSiFloor(c *fiber.Ctx) error {
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
  }

  // Start the server, else log errors
  log.Fatal(app.Listen(":" + PORT))
  fmt.Println("Server running on port 4000")
}
