package main

import (
	"log"

	"short-url-api/internal/db"
	"short-url-api/internal/handler"
	"short-url-api/internal/repository"
	"short-url-api/internal/routes"
	"short-url-api/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	loadEnv()

	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	linkRepo := repository.NewSQLLinkRepo(db.DB)

	linkService := service.NewLinkService(linkRepo)

	linkHandler := handler.NewLinkHandler(linkService)

	r := gin.Default()
	r.Use(cors.Default())

	// Register routes
	routes.RegisterRoutes(
		r,
		routes.NewLinkRoutes(linkHandler),
	)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}