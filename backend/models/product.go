package models

type Product struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       string     `json:"price"`
	ImageURL    string     `json:"imageUrl"` // JSON tag disesuaikan dengan camelCase frontend
	Features    StringList `gorm:"type:json" json:"features"`
	BuyURL      string     `json:"buyUrl"`
	Tag         string     `json:"tag"`
}
