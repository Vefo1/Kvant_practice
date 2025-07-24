package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Vefo1/Kvant_practice/config"
	"github.com/Vefo1/Kvant_practice/internal/handler"
	"github.com/Vefo1/Kvant_practice/internal/middleware"
	"github.com/Vefo1/Kvant_practice/internal/services"
	"github.com/Vefo1/Kvant_practice/pkg/logger"
)

func main() {
	// 1. Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		// Use standard log for fatal errors before custom logger is initialized
		fmt.Fprintf(os.Stderr, "Fatal error loading configuration: %v\n", err)
		os.Exit(1)
	}

	// 2. Initialize logger
	log := logger.NewLogger(cfg.Logging.Level)
	log.Info("Configuration loaded successfully. Log level: %s", cfg.Logging.Level)

	// 3. Initialize service
	// Pass BaseURL instead of URL
	predictService := services.NewPredictService(cfg.ExternalAPI.BaseURL, cfg.ExternalAPI.Token, log)
	log.Info("Prediction service initialized with external API BaseURL: %s", cfg.ExternalAPI.BaseURL)

	// 4. Initialize handler
	h := handler.NewHandler(predictService, log)
	log.Info("HTTP handler initialized.")

	// 5. Apply middleware
	authMiddleware := middleware.AuthMiddleware(cfg.AppAuth.Token, cfg.AppAuth.HeaderName, log)

	// 6. Register handlers with middleware for each prediction type
	http.Handle("/predict/hba1c", authMiddleware(http.HandlerFunc(h.HBA1CPredictHandler)))
	http.Handle("/predict/ldll", authMiddleware(http.HandlerFunc(h.LdllPredictHandler)))
	http.Handle("/predict/ferr", authMiddleware(http.HandlerFunc(h.FerrPredictHandler)))
	http.Handle("/predict/ldl", authMiddleware(http.HandlerFunc(h.LdlPredictHandler)))
	http.Handle("/predict/tg", authMiddleware(http.HandlerFunc(h.TgPredictHandler)))
	http.Handle("/predict/hdl", authMiddleware(http.HandlerFunc(h.HdlPredictHandler)))
	log.Info("All prediction routes registered with authentication middleware.")

	// 7. Start the HTTP server
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Info("Server starting on %s", addr)
	log.Fatal("Server failed to start: %v", http.ListenAndServe(addr, nil))
}
