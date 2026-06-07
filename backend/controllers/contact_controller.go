package controllers

import (
	"net/http"
	"time"

	"github.com/TubagusAldiMY/portofolio-backend/config"
	"github.com/TubagusAldiMY/portofolio-backend/models"
	"github.com/gin-gonic/gin"
)

type sendMessageInput struct {
	Name    string `json:"name"    binding:"required,max=100"`
	Email   string `json:"email"   binding:"required,email,max=200"`
	Content string `json:"content" binding:"required,max=2000"`
}

// POST /api/contact
func SendMessage(c *gin.Context) {
	var input sendMessageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondBindError(c, "contact send", err)
		return
	}

	msg := models.Message{
		Name:      input.Name,
		Email:     input.Email,
		Content:   input.Content,
		CreatedAt: time.Now(),
	}

	if err := config.DB.Create(&msg).Error; err != nil {
		respondInternalError(c, "create contact message", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pesan berhasil dikirim!", "data": msg})
}

// GET /api/messages (Hanya untuk Admin nanti)
func GetMessages(c *gin.Context) {
	var messages []models.Message
	if err := config.DB.Order("created_at desc").Find(&messages).Error; err != nil {
		respondInternalError(c, "list contact messages", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": messages})
}
