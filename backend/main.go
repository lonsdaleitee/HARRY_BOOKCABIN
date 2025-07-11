package main

import (
	"log"
	"net/http"

	"airline-voucher-backend/config"
	"airline-voucher-backend/handlers"
	"airline-voucher-backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.NewConfig()

	// Initialize database
	db, err := config.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize services
	voucherService := services.NewVoucherService(db)

	// Initialize handlers
	voucherHandler := handlers.NewVoucherHandler(voucherService)

	// Initialize Gin router
	router := gin.Default()

	// Configure CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"} // Frontend URL
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(corsConfig))

	// Routes
	api := router.Group("/api")
	{
		api.POST("/check", voucherHandler.CheckVoucher)
		api.POST("/generate", voucherHandler.GenerateVoucher)
	}

	// Health check endpoint
	router.GET("/health", voucherHandler.HealthCheck)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	log.Printf("Database file: %s", cfg.DBPath)
	log.Printf("CORS enabled for: http://localhost:3000")

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
