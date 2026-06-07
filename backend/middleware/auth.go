package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// 1. Ambil token dari header Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	// Format header harus "Bearer <token>"
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	tokenString = strings.TrimSpace(tokenString)
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
		return
	}

	// Guard: Pastikan JWT_SECRET tersedia
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server misconfiguration: JWT secret not set"})
		return
	}

	// 2. Parse dan Validasi Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}

	// 3. Cek Expiration
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if expVal, ok := claims["exp"].(float64); ok {
			if float64(time.Now().Unix()) > expVal {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
				return
			}
		}
		// Simpan user ID ke context (opsional)
		if sub, ok := claims["sub"]; ok {
			c.Set("userID", sub)
		}
	}

	c.Next()
}
