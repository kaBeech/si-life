package main
 
import (
  "fmt"

  "net/http"

  "github.com/gin-gonic/gin"
)
 
func main() {
  router := gin.Default()
 
  // Create API route group
  api := router.Group("/api")
  {
    // Add routes and define route handler functions
    api.GET("/si-life", func(ctx *gin.Context) {
      ctx.JSON(200, gin.H{"msg": "Welcome to Si-Life!"})
    })
  }
 
  // Handle 404s
  router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })
 
  // Start the server
  router.Run(":8080")
  fmt.Println("Server running on port " + port)
}
