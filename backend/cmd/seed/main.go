// Seed script — run once to populate initial project data.
// Usage: go run ./cmd/seed
// Safe to run multiple times; skips entries where title already exists.
package main

import (
	"fmt"
	"log"

	"github.com/TubagusAldiMY/portofolio-backend/config"
	"github.com/TubagusAldiMY/portofolio-backend/models"
)

var products = []models.Product{
	{
		Title: "Backend Architecture Review",
		Description: "Deep architecture review untuk backend service Go/Rust. Mencakup audit Clean Architecture, " +
			"identifikasi anti-pattern, evaluasi inter-service communication (gRPC/REST/message broker), " +
			"strategi distributed data, dan rekomendasi konkret berbasis OWASP Top 10. Output: laporan tertulis + " +
			"session diskusi 60 menit.",
		Price:    "Mulai Rp 2.500.000",
		ImageURL: "",
		Features: models.StringList{
			"Audit Clean Architecture & SOLID",
			"Identifikasi coupling antar service",
			"Review OWASP Top 10 compliance",
			"Rekomendasi refactor prioritas",
			"Sesi diskusi 60 menit",
		},
		BuyURL: "",
		Tag:    "Consulting",
	},
	{
		Title: "Production-Ready Microservice Scaffold",
		Description: "Setup microservice production-ready dari nol: Go + Gin/gRPC, observability stack " +
			"(Prometheus, Grafana, Jaeger), CI/CD pipeline, Docker + docker-compose, structured logging JSON, " +
			"JWT auth, rate limiting, graceful shutdown, dan integration test. Cocok untuk founder yang butuh " +
			"fondasi backend solid tanpa harus mulai dari blank repo.",
		Price:    "Mulai Rp 7.500.000",
		ImageURL: "",
		Features: models.StringList{
			"Go + Gin/gRPC base",
			"Observability stack",
			"CI/CD GitHub Actions",
			"Docker compose dev environment",
			"Auth + rate limit + graceful shutdown",
		},
		BuyURL: "",
		Tag:    "Engineering",
	},
}

var projects = []models.Project{
	{
		Title: "KasKu — Financial Management SaaS",
		Description: "Production-ready SaaS platform for personal and small-business finance, " +
			"built as a microservices monorepo. Features a multi-tenant offline-first PWA frontend " +
			"with real-time sync, comprehensive financial dashboard (accounts, transactions, budgets, investments), " +
			"subscription/payment orchestration, and an enterprise observability stack. " +
			"Demonstrates event-driven async flows, gRPC inter-service communication, and database-per-service isolation.",
		TechStack: models.StringList{
			"Go", "Rust", "SvelteKit", "Svelte 5", "TypeScript",
			"PostgreSQL", "Redis", "RabbitMQ", "gRPC", "Docker",
			"Prometheus", "Grafana", "Jaeger",
		},
		RepoUrl: "https://github.com/TubagusAldiMY/kasku",
		LiveUrl: "",
	},
	{
		Title: "Music Player — Native Desktop App",
		Description: "Native Ubuntu desktop music player built with Tauri 2 and Rust. " +
			"Supports offline playback of high-quality audio formats including FLAC, " +
			"local library management with incremental scanning, synced lyrics (LRC sidecar and embedded metadata), " +
			"and seamless desktop integration via MPRIS and media key controls. " +
			"Clean architecture separation across Rust workspace layers with headless playback logic that stays framework-independent.",
		TechStack: models.StringList{
			"Tauri 2", "Rust", "React", "TypeScript",
			"TailwindCSS", "SQLite", "Axum", "rodio", "tokio", "Bun",
		},
		RepoUrl: "https://github.com/TubagusAldiMY/music-player",
		LiveUrl: "",
	},
}

func main() {
	config.ConnectDatabase()

	inserted := 0
	skipped := 0

	fmt.Println("Projects:")
	for _, p := range projects {
		var existing models.Project
		if err := config.DB.Where("title = ?", p.Title).First(&existing).Error; err == nil {
			fmt.Printf("  SKIP  %s\n", p.Title)
			skipped++
			continue
		}
		if result := config.DB.Create(&p); result.Error != nil {
			log.Printf("  ERROR %s: %v\n", p.Title, result.Error)
			continue
		}
		fmt.Printf("  OK    %s (id=%d)\n", p.Title, p.ID)
		inserted++
	}

	fmt.Println("\nProducts:")
	for _, p := range products {
		var existing models.Product
		if err := config.DB.Where("title = ?", p.Title).First(&existing).Error; err == nil {
			fmt.Printf("  SKIP  %s\n", p.Title)
			skipped++
			continue
		}
		if result := config.DB.Create(&p); result.Error != nil {
			log.Printf("  ERROR %s: %v\n", p.Title, result.Error)
			continue
		}
		fmt.Printf("  OK    %s (id=%d)\n", p.Title, p.ID)
		inserted++
	}

	fmt.Printf("\nDone: %d inserted, %d skipped.\n", inserted, skipped)
}
