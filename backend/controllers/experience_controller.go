package controllers

import (
	"net/http"

	"github.com/TubagusAldiMY/portofolio-backend/config"
	"github.com/TubagusAldiMY/portofolio-backend/models"
	"github.com/gin-gonic/gin"
)

type experienceInput struct {
	Role        string `json:"role" binding:"required,max=160"`
	Company     string `json:"company" binding:"required,max=160"`
	Duration    string `json:"duration" binding:"required,max=120"`
	Description string `json:"description" binding:"required,max=5000"`
}

func (input experienceInput) applyTo(experience *models.Experience) {
	experience.Role = input.Role
	experience.Company = input.Company
	experience.Duration = input.Duration
	experience.Description = input.Description
}

// GET /api/experiences (Public)
func GetExperiences(c *gin.Context) {
	var experiences []models.Experience
	if err := config.DB.Order("id desc").Find(&experiences).Error; err != nil {
		respondInternalError(c, "list experiences", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": experiences})
}

// POST /api/admin/experiences (Admin)
func CreateExperience(c *gin.Context) {
	var input experienceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondBindError(c, "create experience", err)
		return
	}

	var experience models.Experience
	input.applyTo(&experience)

	if err := config.DB.Create(&experience).Error; err != nil {
		respondInternalError(c, "create experience", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": experience})
}

// PUT /api/admin/experiences/:id (Admin)
func UpdateExperience(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	var experience models.Experience
	if err := config.DB.First(&experience, id).Error; err != nil {
		respondLookupError(c, "find experience for update", err, "Experience not found")
		return
	}

	var input experienceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondBindError(c, "update experience", err)
		return
	}

	input.applyTo(&experience)
	if err := config.DB.Save(&experience).Error; err != nil {
		respondInternalError(c, "update experience", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": experience})
}

// DELETE /api/admin/experiences/:id (Admin)
func DeleteExperience(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	var experience models.Experience
	if err := config.DB.First(&experience, id).Error; err != nil {
		respondLookupError(c, "find experience for delete", err, "Experience not found")
		return
	}
	if err := config.DB.Delete(&experience).Error; err != nil {
		respondInternalError(c, "delete experience", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Experience deleted"})
}
