package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"latih.in-be/config"
	"latih.in-be/internal/seeder"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("ENV LOAD ERROR:", err)
	}

	caCertPath := "cert/ca.pem"
	caCert, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("Failed to read CA file: %v", err)
	}

	rootCertPool := x509.NewCertPool()
	if ok := rootCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal("Failed to append CA certificate")
	}

	err = mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: rootCertPool,
	})
	if err != nil {
		log.Fatalf("Failed to register TLS config: %v", err)
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
