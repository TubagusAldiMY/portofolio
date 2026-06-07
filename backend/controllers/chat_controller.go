package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/TubagusAldiMY/portofolio-backend/models"
	"github.com/TubagusAldiMY/portofolio-backend/services"
	"github.com/gin-gonic/gin"
)

const (
	maxChatMessageBytes = 2000
	geminiTimeout       = 15 * time.Second
)

var geminiClient = &http.Client{Timeout: geminiTimeout}

type geminiRequest struct {
	Contents []geminiContent `json:"contents"`
}
type geminiContent struct {
	Role  string       `json:"role"`
	Parts []geminiPart `json:"parts"`
}
type geminiPart struct {
	Text string `json:"text"`
}

type geminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []geminiPart `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func ChatWithAI(c *gin.Context) {
	var req models.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		respondBindError(c, "chat", err)
		return
	}

	if len(req.Message) > maxChatMessageBytes {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message too long (max 2000 characters)"})
		return
	}

	ragContext := services.GetContext()
	prompt := "Kamu adalah asisten AI profesional untuk portofolio Tubagus Aldi. " +
		"Tugasmu adalah menjawab pertanyaan pengunjung HANYA berdasarkan informasi berikut ini. " +
		"Jika informasi tidak ada di konteks, katakan dengan sopan bahwa kamu tidak tahu dan sarankan menghubungi via email.\n\n" +
		ragContext + "\n\nUser: " + req.Message + "\nAssistant:"

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI belum dikonfigurasi"})
		return
	}

	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key=" + apiKey

	body, err := json.Marshal(geminiRequest{
		Contents: []geminiContent{
			{Role: "user", Parts: []geminiPart{{Text: prompt}}},
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare request"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), geminiTimeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare request"})
		return
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := geminiClient.Do(request)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "AI service unavailable"})
		return
	}
	defer resp.Body.Close()

	rawBody, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20)) // cap at 1 MB
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read AI response"})
		return
	}

	if resp.StatusCode >= http.StatusBadRequest {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "AI is busy or limit reached"})
		return
	}

	if text, ok := extractGeminiText(rawBody); ok {
		c.JSON(http.StatusOK, models.ChatResponse{Reply: text})
		return
	}

	c.JSON(http.StatusServiceUnavailable, gin.H{"error": "AI is busy or limit reached"})
}

func extractGeminiText(body []byte) (string, bool) {
	var result geminiResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", false
	}

	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		return "", false
	}

	text := result.Candidates[0].Content.Parts[0].Text
	if text == "" {
		return "", false
	}

	return text, true
}
