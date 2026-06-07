package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/TubagusAldiMY/portofolio-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, continuing with environment variables")
	}

	// Format DSN khusus MySQL: "user:pass@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Ganti postgres.Open menjadi mysql.Open
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err := database.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.Message{},
		&models.Project{},
		&models.Experience{},
	); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	DB = database
	fmt.Println("Database connected & migrated (MySQL)")
}
