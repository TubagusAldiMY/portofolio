package services

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	knowledgeContext string
	once             sync.Once
)

// LoadKnowledgeBase membaca semua file .md/.txt dari folder target
func LoadKnowledgeBase(dirPath string) {
	once.Do(func() {
		var sb strings.Builder
		sb.WriteString("Ini adalah informasi referensi (Knowledge Base) tentang Tubagus Aldi:\n\n")

		err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Hanya baca file markdown atau text
			if !info.IsDir() && (strings.HasSuffix(info.Name(), ".md") || strings.HasSuffix(info.Name(), ".txt")) {
				content, err := os.ReadFile(path)
				if err != nil {
					log.Printf("Gagal membaca file %s: %v", path, err)
					return nil
				}
				fmt.Fprintf(&sb, "--- Sumber: %s ---\n", info.Name())
				sb.Write(content)
				sb.WriteString("\n\n")
			}
			return nil
		})

		if err != nil {
			log.Println("Error loading knowledge base:", err)
		} else {
			knowledgeContext = sb.String()
			fmt.Println("📚 Knowledge Base loaded into memory!")
		}
	})
}

// GetContext mengembalikan string context yang sudah di-cache
func GetContext() string {
	return knowledgeContext
}
