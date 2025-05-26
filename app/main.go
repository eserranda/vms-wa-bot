package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"vms-bot/notification"
	"vms-bot/whatsapp"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	godotenv.Load(".env")
	ctx := context.Background()

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	router.Use(CORSMiddleware())

	wa_client, err := whatsapp.NewWhatsappmeowClient()
	if err != nil {
		log.Fatalf("Failed to create WhatsApp client: %v", err)
		return
	}

	if err := wa_client.Connect(ctx); err != nil {
		log.Fatalf("Failed to connect to WhatsApp: %v", err)
		return
	}
	defer wa_client.Disconnect()

	wa_client.SetEventsHandler(ctx)

	sendNotification := notification.NewNotificationHandler(wa_client)

	notification := router.Group("/api/send-notification")
	sendNotification.SendNotification(notification)

	// Start server
	router.Run(":" + os.Getenv("APP_PORT"))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}
