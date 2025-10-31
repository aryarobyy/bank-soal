package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	model "latih.in-be/internal/model"
)

var databaseInstance *gorm.DB

func InitDB() *gorm.DB {

	var err error
	databaseInstance, err = connectionDatabase()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	err = performMigration()
	if err != nil {
		log.Fatalf("Could not auto migrate: %v", err)
	}

	return databaseInstance

}

func connectionDatabase() (*gorm.DB, error) {

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	databaseConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	sqlDatabase, err := databaseConnection.DB()
	if err != nil {
		return nil, err
	}

	idleConns, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	openConns, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	connLifetime, _ := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME"))

	sqlDatabase.SetMaxIdleConns(idleConns)
	sqlDatabase.SetMaxOpenConns(openConns)
	sqlDatabase.SetConnMaxLifetime(time.Duration(connLifetime) * time.Second)

	return databaseConnection, nil
}

func performMigration() error {
	err := databaseInstance.AutoMigrate(
		&model.User{},
		&model.Question{},
		&model.Option{},
		&model.ExamScore{},
		&model.Exam{},
		&model.ExamSession{},
	)
	if err != nil {
		return err
	}
	return nil
}
