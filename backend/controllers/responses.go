package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func respondInternalError(c *gin.Context, context string, err error) {
	log.Printf("%s: %v", context, err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
}

func respondLookupError(c *gin.Context, context string, err error, notFoundMessage string) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": notFoundMessage})
		return
	}

	respondInternalError(c, context, err)
}

// respondBindError logs the actual binding/validation error (for debugging)
// but returns a generic message to avoid leaking struct field names or
// JSON parser internals to clients.
func respondBindError(c *gin.Context, context string, err error) {
	log.Printf("%s: bind error: %v", context, err)
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
}

// parseIDParam parses a numeric route param (e.g. ":id") as uint. On failure
// it writes a 400 and returns ok=false so callers can return early.
func parseIDParam(c *gin.Context, name string) (uint, bool) {
	raw := c.Param(name)
	id, err := strconv.ParseUint(raw, 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return 0, false
	}
	return uint(id), true
}
