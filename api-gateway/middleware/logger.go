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
	Method     string `json:"method"`
	Path       string `json:"path"`
	Status     int    `json:"status"`
	IP         string `json:"ip"`
	UserAgent  string `json:"user_agent"`
	DurationMs int64  `json:"duration_ms"`
	Error      string `json:"error,omitempty"`
	Timestamp  string `json:"timestamp"`
}

func GlobalLogger() fiber.Handler {
    return func(c *fiber.Ctx) error {
        start := time.Now()

        err := c.Next()
        duration := time.Since(start).Milliseconds()

        logItem := LogModel{
            Method:     c.Method(),
            Path:       c.Path(),
            Status:     c.Response().StatusCode(),
            IP:         c.IP(),
            UserAgent:  string(c.Request().Header.UserAgent()),
            DurationMs: duration,
            Timestamp:  time.Now().Format(time.RFC3339),
        }

        if err != nil {
            logItem.Error = err.Error()
        }

        go sendToElastic(logItem)

        return err
    }
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
