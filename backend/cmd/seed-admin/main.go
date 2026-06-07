// Seed admin user for production bootstrap.
// Usage: ADMIN_USERNAME=admin ADMIN_PASSWORD='long-password' go run ./cmd/seed-admin
package main

import (
	"errors"
	"log"
	"os"

	"github.com/TubagusAldiMY/portofolio-backend/config"
	"github.com/TubagusAldiMY/portofolio-backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")

	if username == "" {
		log.Fatal("ADMIN_USERNAME is required")
	}
	if len(password) < 12 {
		log.Fatal("ADMIN_PASSWORD must be at least 12 characters")
	}

	config.ConnectDatabase()

	var existing models.User
	err := config.DB.Where("username = ?", username).First(&existing).Error
	if err == nil {
		log.Fatalf("Admin user %q already exists; refusing to overwrite", username)
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Fatalf("Failed to check existing admin: %v", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		log.Fatalf("Failed to create admin user: %v", err)
	}

	log.Printf("Admin user %q created", username)
}
