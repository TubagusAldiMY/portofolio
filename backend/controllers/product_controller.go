package controllers

import (
	"net/http"

	"github.com/TubagusAldiMY/portofolio-backend/config"
	"github.com/TubagusAldiMY/portofolio-backend/models"
	"github.com/gin-gonic/gin"
)

type productInput struct {
	Title       string            `json:"title" binding:"required,max=160"`
	Description string            `json:"description" binding:"required,max=5000"`
	Price       string            `json:"price" binding:"max=80"`
	ImageURL    string            `json:"imageUrl" binding:"max=2048"`
	Features    models.StringList `json:"features" binding:"max=80,dive,max=160"`
	BuyURL      string            `json:"buyUrl" binding:"omitempty,url,max=2048"`
	Tag         string            `json:"tag" binding:"max=80"`
}

func (input productInput) applyTo(product *models.Product) {
	product.Title = input.Title
	product.Description = input.Description
	product.Price = input.Price
	product.ImageURL = input.ImageURL
	product.Features = input.Features
	product.BuyURL = input.BuyURL
	product.Tag = input.Tag
}

// GET /api/products
func GetProducts(c *gin.Context) {
	var products []models.Product

	if err := config.DB.Find(&products).Error; err != nil {
		respondInternalError(c, "list products", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// POST /api/admin/products
func CreateProduct(c *gin.Context) {
	var input productInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondBindError(c, "create product", err)
		return
	}

	var product models.Product
	input.applyTo(&product)

	if err := config.DB.Create(&product).Error; err != nil {
		respondInternalError(c, "create product", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": product})
}

// PUT /api/admin/products/:id
func UpdateProduct(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		respondLookupError(c, "find product for update", err, "Product not found")
		return
	}

	var input productInput
	if err := c.ShouldBindJSON(&input); err != nil {
		respondBindError(c, "update product", err)
		return
	}

	input.applyTo(&product)
	if err := config.DB.Save(&product).Error; err != nil {
		respondInternalError(c, "update product", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DELETE /api/admin/products/:id
func DeleteProduct(c *gin.Context) {
	id, ok := parseIDParam(c, "id")
	if !ok {
		return
	}

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		respondLookupError(c, "find product for delete", err, "Product not found")
		return
	}
	if err := config.DB.Delete(&product).Error; err != nil {
		respondInternalError(c, "delete product", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
