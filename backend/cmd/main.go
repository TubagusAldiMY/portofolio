package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/TubagusAldiMY/portofolio-backend/config"
	"github.com/TubagusAldiMY/portofolio-backend/controllers"
	"github.com/TubagusAldiMY/portofolio-backend/middleware"
	"github.com/TubagusAldiMY/portofolio-backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	defaultFrontendOrigins = "http://localhost:5173,http://127.0.0.1:5173,http://localhost:5174,http://127.0.0.1:5174,http://localhost:5175,http://127.0.0.1:5175"
	jsonBodyLimitBytes     = 1 << 20
	uploadBodyLimitBytes   = 12 << 20
)

func main() {
	config.ConnectDatabase()
	services.LoadKnowledgeBase("./knowledge_base")

	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET is not set — refusing to start")
	}

	if strings.EqualFold(os.Getenv("APP_ENV"), "production") && os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.MaxMultipartMemory = 10 << 20

	r.Use(limitRequestBody)

	// CORS
	allowedOrigins, allowAllOrigins := parseFrontendOrigins()
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	if allowAllOrigins {
		corsConfig.AllowAllOrigins = true
		corsConfig.AllowCredentials = false
	} else {
		corsConfig.AllowOrigins = allowedOrigins
	}
	r.Use(cors.New(corsConfig))

	// Static uploads
	uploadDir := os.Getenv("UPLOAD_DIR")
	if uploadDir == "" {
		uploadDir = "./public/uploads"
	}
	r.Static("/uploads", uploadDir)

	// Health check (used by load balancers and orchestrators)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api")
	{
		loginLimiter := middleware.RateLimitPerMinute(10, 10)
		contactLimiter := middleware.RateLimitPerMinute(3, 3)
		chatLimiter := middleware.RateLimitPerMinute(5, 5)

		api.POST("/auth/login", loginLimiter, controllers.Login)
		if os.Getenv("ALLOW_REGISTER") == "true" {
			api.POST("/auth/register", loginLimiter, controllers.RegisterAdmin)
		}
		api.POST("/contact", contactLimiter, controllers.SendMessage)
		api.GET("/products", controllers.GetProducts)
		api.GET("/projects", controllers.GetProjects)
		api.GET("/projects/:id", controllers.GetProject)
		api.GET("/experiences", controllers.GetExperiences)
		api.POST("/chat", chatLimiter, controllers.ChatWithAI)

		admin := api.Group("/admin")
		admin.Use(middleware.RequireAuth)
		{
			admin.GET("/messages", controllers.GetMessages)
			admin.POST("/upload", controllers.UploadFile)
			admin.POST("/projects", controllers.CreateProject)
			admin.PUT("/projects/:id", controllers.UpdateProject)
			admin.DELETE("/projects/:id", controllers.DeleteProject)
			admin.POST("/products", controllers.CreateProduct)
			admin.PUT("/products/:id", controllers.UpdateProduct)
			admin.DELETE("/products/:id", controllers.DeleteProduct)
			admin.POST("/experiences", controllers.CreateExperience)
			admin.PUT("/experiences/:id", controllers.UpdateExperience)
			admin.DELETE("/experiences/:id", controllers.DeleteExperience)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Forced shutdown: %v", err)
		}
	}()

	log.Printf("Starting server on :%s", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
	log.Println("Server stopped")
}

func limitRequestBody(c *gin.Context) {
	limit := int64(jsonBodyLimitBytes)
	if strings.HasPrefix(strings.ToLower(c.GetHeader("Content-Type")), "multipart/form-data") {
		limit = uploadBodyLimitBytes
	}

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, limit)
	c.Next()
}

func parseFrontendOrigins() ([]string, bool) {
	originsEnv := strings.TrimSpace(os.Getenv("FRONTEND_ORIGINS"))
	if originsEnv == "" {
		originsEnv = defaultFrontendOrigins
	}
	if originsEnv == "*" {
		return nil, true
	}

	allowedOrigins := make([]string, 0)
	for _, part := range strings.Split(originsEnv, ",") {
		origin := strings.TrimSpace(part)
		if origin != "" {
			allowedOrigins = append(allowedOrigins, origin)
		}
	}
	if len(allowedOrigins) == 0 {
		return strings.Split(defaultFrontendOrigins, ","), false
	}

	return allowedOrigins, false
}
