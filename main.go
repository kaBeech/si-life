package main
 
import (
  "fmt"
  "log"

  "github.com/gofiber/fiber/v2"
)
 
func main() {
  app := fiber.New()
 // 
 //  // Create API route group
 //  api := router.Group("/api")
 //  {
 //    // Add routes and define route handler functions
 //    api.GET("/si-life", func(ctx *gin.Context) {
 //      ctx.JSON(200, gin.H{"msg": "Welcome to Si-Life!"})
 //    })
 //  }
 // 
 //  // Handle 404s
 //  router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })
 
  // Start the server, else log errors
  log.Fatal(app.Listen(":4000"))
  fmt.Println("Server running on port 4000")
}
