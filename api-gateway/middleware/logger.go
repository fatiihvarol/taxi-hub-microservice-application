package middleware

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"api-gateway/elastic"
)

type LogModel struct {
	Method      string `json:"method"`
	Path        string `json:"path"`
	Status      int    `json:"status"`
	IP          string `json:"ip"`
	UserAgent   string `json:"user_agent"`
	DurationMs  int64  `json:"duration_ms"`
	RequestBody string `json:"request_body,omitempty"`
	Response    string `json:"response_body,omitempty"`
	Error       string `json:"error,omitempty"`
	Timestamp   string `json:"timestamp"`
}

func GlobalLogger(c *fiber.Ctx) error {
	start := time.Now()

	// Request body yakala
	reqBody := string(c.Body())

	// Handler çalıştır
	err := c.Next()
	duration := time.Since(start).Milliseconds()

	// Response body yakala
	resBody := string(c.Response().Body())

	logItem := LogModel{
		Method:      c.Method(),
		Path:        c.Path(),
		Status:      c.Response().StatusCode(),
		IP:          c.IP(),
		UserAgent:   string(c.Request().Header.UserAgent()),
		DurationMs:  duration,
		RequestBody: reqBody,
		Response:    resBody,
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	// Hata varsa ekle
	if err != nil {
		logItem.Error = err.Error()
	}

	// Elastic'e async gönder
	go sendToElastic(logItem)

	return err
}

func sendToElastic(item LogModel) {
	data, _ := json.Marshal(item)

	res, err := elastic.ES.Index(
		elastic.ElasticIndex,
		bytes.NewReader(data),
	)
	if err != nil {
		log.Printf("[Elastic] insert error: %v", err)
		return
	}
	defer res.Body.Close()
}
