package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
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

	caCertBase64 := os.Getenv("DB_CA_CERT_BASE64")

	var caCert []byte

	if caCertBase64 != "" {
		// Mode 1: Pakai ENV base64
		decoded, err := base64.StdEncoding.DecodeString(caCertBase64)
		if err != nil {
			log.Fatalf("Failed to decode DB_CA_CERT_BASE64: %v", err)
		}
		caCert = decoded
		log.Println("Loaded CA cert from ENV (base64)")
	} else {
		// Mode 2: Pakai file lokal
		localPath := "./cert/ca.pem"

		fileBytes, err := os.ReadFile(localPath)
		if err != nil {
			log.Fatalf("DB_CA_CERT_BASE64 is empty and failed to read local CA file: %v", err)
		}

		caCert = fileBytes
		log.Println("Loaded CA cert from local file:", localPath)
	}

	rootCertPool := x509.NewCertPool()
	if !rootCertPool.AppendCertsFromPEM(caCert) {
		log.Fatal("Failed to append CA certificate")
	}

	err := mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: rootCertPool,
	})
	if err != nil {
		log.Fatalf("Failed to register TLS config: %v", err)
	}

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
	if err := seeder.SeedUser(db); err != nil {
		fmt.Println("Failed seed user:", err)
	} else {
		fmt.Println("User seeding success")
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
