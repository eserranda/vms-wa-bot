package notification

import (
	"context"
	"fmt"
	"vms-bot/model"
	"vms-bot/utils"

	"github.com/gin-gonic/gin"
	"go.mau.fi/whatsmeow/types"
)

type NotificationHandler struct {
	wa model.WhatsAppClient
}

func NewNotificationHandler(wa model.WhatsAppClient) *NotificationHandler {
	return &NotificationHandler{
		wa: wa,
	}
}

func (h *NotificationHandler) SendNotification(r *gin.RouterGroup) {
	r.POST("/message", h.TamuMasuk)
	r.POST("/guest", h.GuestEntry)
	r.POST("/guest-check-in", h.GuestCheckIn)
}

func (h *NotificationHandler) GuestCheckIn(c *gin.Context) {
	var request struct {
		UserPhoneNumber string `json:"user_phone_number"`
		UserName        string `json:"user_name"`
		SecurityName    string `json:"security_name"`

		VisitorName string `json:"visitor_name"`
		CheckIn     string `json:"check_in"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error ya": err.Error()})
		return
	}

	fmt.Println("Request: ", request)
	infoChat := types.JID{
		User:   request.UserPhoneNumber,
		Server: "s.whatsapp.net",
	}

	greeting := utils.GetGreetingBasedOnTime()
	message := ""

	message += fmt.Sprintf("%s *%s*\n\n", greeting, request.UserName)
	message += fmt.Sprintf("%s *%s* %s *%s* %s\n\n", "Tamu atas nama", request.VisitorName, "telah di verifikasi oleh", request.SecurityName, "(security)")
	message += "Berikut detail kunjungan:\n\n"
	message += fmt.Sprintf("Nama Tamu : *%s*\n", request.VisitorName)
	message += fmt.Sprintf("Waktu : *%s*\n", request.CheckIn)
	message += fmt.Sprintf("Catatan :\n*%s*\n\n", "Tamu telah melakukan check-in.")
	message += "_Visitor Management System_\n"
	message += "_powered by VMS Bot Notification_"

	if err := h.wa.SendMessage(context.Background(), infoChat, types.JID{}, message); err != nil {
		fmt.Println("Error : ", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Notification sent successfully",
	})
}

func (h *NotificationHandler) GuestEntry(c *gin.Context) {
	var request struct {
		UserPhoneNumber string `json:"user_phone_number"`
		UserName        string `json:"user_name"`
		SecurityName    string `json:"security_name"`

		VisitorName        string `json:"visitor_name"`
		VisitorPhoneNumber string `json:"visitor_phone_number"`
		VisitorType        string `json:"visitor_type"`
		CompanyName        string `json:"company_name"`
		VehicleNumber      string `json:"vehicle_number"`
		Purpose            string `json:"purpose"`
		CheckIn            string `json:"check_in"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error ya": err.Error()})
		return
	}

	infoChat := types.JID{
		User:   request.UserPhoneNumber,
		Server: "s.whatsapp.net",
	}

	greeting := utils.GetGreetingBasedOnTime()
	message := ""

	message += fmt.Sprintf("%s *%s*\n\n", greeting, request.UserName)
	message += fmt.Sprintf("%s *%s* %s *%s* %s\n\n", "Tamu atas nama", request.VisitorName, "telah di verifikasi oleh", request.SecurityName, "(security) untuk mengunjungi Anda.")
	message += "Berikut detail kunjungan:\n\n"
	message += fmt.Sprintf("Nama Tamu : *%s*\n", request.VisitorName)
	message += fmt.Sprintf("Pengunjung : *%s* (%s)\n", request.VisitorType, request.CompanyName)
	message += fmt.Sprintf("No. Plat : *%s*\n", request.VehicleNumber)
	message += fmt.Sprintf("Waktu : *%s*\n", request.CheckIn)
	message += fmt.Sprintf("Tujuan Berkunjung :\n *%s*\n\n", request.Purpose)
	message += "_Visitor Management System_\n"
	message += "_powered by VMS Bot Notification_"

	if err := h.wa.SendMessage(context.Background(), infoChat, types.JID{}, message); err != nil {
		fmt.Println("Error : ", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Notification sent successfully",
	})
}

func (h *NotificationHandler) TamuMasuk(c *gin.Context) {

	var request struct {
		// VisitorNumber   string `json:"visitor_number"`
		UserPhoneNumber    string `json:"user_phone_number"`
		UserName           string `json:"user_name"`
		VisitorName        string `json:"visitor_name"`
		VisitorPhoneNumber string `json:"visitor_phone_number"`
		ArrivalDate        string `json:"arrival_date"`
		ArrivalTime        string `json:"arrival_time"`
		VerificationURI    string `json:"verification_uri"`
		Message            string `json:"message"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	infoChat := types.JID{
		User:   request.UserPhoneNumber,
		Server: "s.whatsapp.net",
	}

	greeting := utils.GetGreetingBasedOnTime()
	message := ""

	// message = "\n\n"
	message += fmt.Sprintf("%s *%s*\n\n", greeting, request.UserName)
	message += fmt.Sprintf("%s *%s*%s\n\n", "Tamu atas nama", request.VisitorName, " baru saja melakukan registrasi untuk mengunjungi Anda.")
	message += "Berikut detail kunjungan:\n\n"
	message += fmt.Sprintf("Nama Tamu : *%s*\n", request.VisitorName)
	message += fmt.Sprintf("No. HP : *%s*\n", request.VisitorPhoneNumber)
	message += fmt.Sprintf("Tanggal Berkunjung: *%s*\n", request.ArrivalDate)
	message += fmt.Sprintf("Jam (estimasi): *%s*\n\n", request.ArrivalTime)
	message += fmt.Sprintf("Silakan klik link dibawah ini untuk melakukan verifikasi:\n%s\n\n", request.VerificationURI)
	message += "_Visitor Management System_\n"
	message += "_powered by VMS Bot Notification_"

	if err := h.wa.SendMessage(context.Background(), infoChat, types.JID{}, message); err != nil {
		fmt.Println("Error : ", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Notification sent successfully",
	})
}
