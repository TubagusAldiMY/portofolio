package models

import "time"

// Model untuk Admin Login
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"-"` // Password tidak dikirim balik di JSON
}

// Model untuk Pesan Masuk (Contact Form)
type Message struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Content   string    `json:"content"` // Isi pesan
	CreatedAt time.Time `json:"created_at"`
}

// Model untuk Projects (Portofolio)
type Project struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ImageUrl    string     `json:"imageUrl"`
	TechStack   StringList `gorm:"type:json" json:"techStack"`
	RepoUrl     string     `json:"repoUrl"`
	LiveUrl     string     `json:"liveUrl"`
}

// Model untuk Experience
type Experience struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Role        string `json:"role"`
	Company     string `json:"company"`
	Duration    string `json:"duration"`
	Description string `json:"description"` // Bisa disimpan sebagai JSON text jika panjang
}
