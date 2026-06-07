package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

const maxUploadBytes = 10 << 20 // 10 MB

var allowedMIME = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
	"image/gif":  true,
}

func randomHex(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// POST /api/admin/upload
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	if file.Size > maxUploadBytes {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large (max 10 MB)"})
		return
	}

	opened, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read uploaded file"})
		return
	}
	defer opened.Close()

	mt, err := mimetype.DetectReader(opened)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to detect file type"})
		return
	}
	if !allowedMIME[mt.String()] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported file type"})
		return
	}

	uploadPath := os.Getenv("UPLOAD_DIR")
	if uploadPath == "" {
		uploadPath = "./public/uploads"
	}
	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	id, err := randomHex(16)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate file name"})
		return
	}
	filename := id + mt.Extension()
	dest := filepath.Join(uploadPath, filename)

	if err := c.SaveUploadedFile(file, dest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	appURL := os.Getenv("APP_URL")
	relPath := "/uploads/" + filename
	fileURL := relPath
	if appURL != "" {
		fileURL = appURL + relPath
	}

	c.JSON(http.StatusOK, gin.H{"url": fileURL, "filename": filename})
}
