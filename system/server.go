package system

import (
	"log"
	"my-shortener/controller"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartServer() {
	// Initialize your server and routes here.
	// For example, you can use the net/http package to create a simple HTTP server.
	godotenv.Load()

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := gin.Default() // ← เปลี่ยนจาก gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.SetTrustedProxies(nil)
	router.GET("/ping", ping)
	controller.UrlController(router)

	// Start the server on port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...", port)
	router.Run(":" + port)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
