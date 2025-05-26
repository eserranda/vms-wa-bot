package utils

import (
	"log"
	"os"
	"time"
)

func GetGreetingBasedOnTime() string {
	timeZone := os.Getenv("TIME_ZONE")
	if timeZone == "" {
		timeZone = "Asia/Jakarta"
	}

	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		log.Println("Error loading time zone:", err)
		return "Hello!"
	}

	// ambil waktu saat ini berdasarkan zona waktu
	currentTime := time.Now().In(loc)

	// mendapatkan jam saat ini
	hour := currentTime.Hour()

	// menemukan salam berdasarkan jam
	switch {
	case hour >= 4 && hour < 10:
		return "Selamat Pagi"
	case hour >= 10 && hour < 14:
		return "Selamat Siang"
	case hour >= 14 && hour < 18:
		return "Selamat Sore"
	default:
		return "Selamat Malam"
	}
}
