package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"latih.in-be/config"
	"latih.in-be/internal/seeder"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	db := config.InitDB()
	application := config.NewApp(db)

	application.Router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Hello world"})
	})

	if err := seeder.SeedSubjects(db); err != nil {
		fmt.Println("Failed seed subjects:", err)
	} else {
		fmt.Println("Subject seeding success")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s with CORS enabled", port)

	if err := application.Run(":" + port); err != nil {
		fmt.Println("FATAL ERROR: Server failed to run:", err)
		os.Exit(1)
	}
}
