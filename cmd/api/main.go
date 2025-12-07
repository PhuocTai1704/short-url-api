package main

import (
	"log"
	"os"

	"short-url-api/internal/db"
	"short-url-api/internal/handler"
	"short-url-api/internal/repository"
	"short-url-api/internal/routes"
	"short-url-api/internal/service"

	_ "short-url-api/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	loadEnv()

	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Tạo các repo
	linkRepo := repository.NewSQLLinkRepo(db.DB)
	linkClickRepo := repository.NewSQLLinkClickRepo(db.DB)

	// Tạo service cho LinkClick trước
	linkClickService := service.NewLinkClickService(linkClickRepo)

	// Tạo service cho Link, truyền LinkRepo và LinkClickService
	linkService := service.NewLinkService(linkRepo, linkClickService)

	// Tạo handler
	linkHandler := handler.NewLinkHandler(linkService)

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes
	routes.RegisterRoutes(
		r,
		routes.NewLinkRoutes(linkHandler),
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
