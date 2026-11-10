package main

import (
	"log"
	"os"

	"backend/config"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" 
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Notice: .env file not found, using system environment variables")
	}

	config.ConnectDB()

	app := fiber.New(fiber.Config{
		AppName: "Feature Management API v1.0",
	})


	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", 
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	app.Use(logger.New())

	api := app.Group("/api/v1")
	routes.SetupFeatureRoutes(api) 

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server is running on port %s...", port)
	log.Fatal(app.Listen(":" + port))
}