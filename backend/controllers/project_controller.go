package controllers

import (
	"net/http"

	"github.com/TubagusAldiMY/portofolio-backend/config"
	"github.com/TubagusAldiMY/portofolio-backend/models"
	"github.com/gin-gonic/gin"
)

type projectInput struct {
	Title       string            `json:"title" binding:"required,max=160"`
	Description string            `json:"description" binding:"required,max=5000"`
	ImageUrl    string            `json:"imageUrl" binding:"max=2048"`
	TechStack   models.StringList `json:"techStack" binding:"max=40,dive,max=80"`
	RepoUrl     string            `json:"repoUrl" binding:"omitempty,url,max=2048"`
	LiveUrl     string            `json:"liveUrl" binding:"omitempty,url,max=2048"`
}

func (input projectInput) applyTo(project *models.Project) {
	project.Title = input.Title
	project.Description = input.Description
	project.ImageUrl = input.ImageUrl
	project.TechStack = input.TechStack
	project.RepoUrl = input.RepoUrl
	project.LiveUrl = input.LiveUrl
}

// GET /api/projects (Public)
func GetProjects(c *gin.Context) {
	var projects []models.Project
	if err := config.DB.Find(&projects).Error; err != nil {
		respondInternalError(c, "list projects", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

// GET /api/projects/:id (Public)
func GetProject(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	var project models.Project
	if err := config.DB.First(&project, id).Error; err != nil {
		respondLookupError(c, "get project", err, "Project not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

// POST /api/admin/projects (Admin)
func CreateProject(c *gin.Context) {
	var input projectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondBindError(c, "create project", err)
		return
	}

	var project models.Project
	input.applyTo(&project)

	if err := config.DB.Create(&project).Error; err != nil {
		respondInternalError(c, "create project", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": project})
}

// PUT /api/admin/projects/:id (Admin)
func UpdateProject(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	var project models.Project
	if err := config.DB.First(&project, id).Error; err != nil {
		respondLookupError(c, "find project for update", err, "Project not found")
		return
	}

	var input projectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondBindError(c, "update project", err)
		return
	}

	input.applyTo(&project)
	if err := config.DB.Save(&project).Error; err != nil {
		respondInternalError(c, "update project", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

// DELETE /api/admin/projects/:id (Admin)
func DeleteProject(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	var project models.Project
	if err := config.DB.First(&project, id).Error; err != nil {
		respondLookupError(c, "find project for delete", err, "Project not found")
		return
	}
	if err := config.DB.Delete(&project).Error; err != nil {
		respondInternalError(c, "delete project", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}
