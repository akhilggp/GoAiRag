package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/akhilggp/GoAiRag/internal/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func checkGET(url string, logger *zap.Logger) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, "", err
	}
	// defer resp.Body.Close()

	defer func() {
		if err := resp.Body.Close(); err != nil {
			// log it (use your logger of choice)
			//.Printf("close response body: %v", err)

			logger.Error("close response body", zap.Error(err))
		}
	}()

	return resp.StatusCode, resp.Status, nil
}

func main() {
	logger, _ := logging.InitZap()
	defer logger.Sync() // flushes buffer, if any

	// defaults match local setup
	ollamaURL := env("OLLAMA_URL", "http://localhost:11434")
	qdrantURL := env("QDRANT_URL", "http://localhost:6333")
	phoenixURL := env("PHOENIX_URL", "http://localhost:6006")
	port := env("PORT", "8080")

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/deps", func(c *gin.Context) {
		// Ollama: API is served under /api (version endpoint works well)
		ollamaCode, ollamaStatus, ollamaErr := checkGET(ollamaURL+"/api/version", logger) // base URL docs mention /api :contentReference[oaicite:3]{index=3}

		// Qdrant: /healthz returns 200 when ready :contentReference[oaicite:4]{index=4}
		qdrantCode, qdrantStatus, qdrantErr := checkGET(qdrantURL+"/healthz", logger)

		// Phoenix: UI served at / on port 6006 :contentReference[oaicite:5]{index=5}
		phoenixCode, phoenixStatus, phoenixErr := checkGET(phoenixURL+"/", logger)

		c.JSON(200, gin.H{
			"ollama": gin.H{
				"url":        ollamaURL,
				"ok":         ollamaErr == nil && ollamaCode >= 200 && ollamaCode < 300,
				"statusCode": ollamaCode,
				"status":     ollamaStatus,
				"error":      errString(ollamaErr),
			},
			"qdrant": gin.H{
				"url":        qdrantURL,
				"ok":         qdrantErr == nil && qdrantCode >= 200 && qdrantCode < 300,
				"statusCode": qdrantCode,
				"status":     qdrantStatus,
				"error":      errString(qdrantErr),
			},
			"phoenix": gin.H{
				"url":        phoenixURL,
				"ok":         phoenixErr == nil && phoenixCode >= 200 && phoenixCode < 400,
				"statusCode": phoenixCode,
				"status":     phoenixStatus,
				"error":      errString(phoenixErr),
			},
		})
	})
	logger.Info("starting server", zap.String("port", port))

	_ = r.Run(":" + port)
}

func errString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
