package server

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"tess-personal/internal/web"
)

func handleContact() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()
		if err != nil {
			web.ContactError().Render(r.Context(), w)
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		tripType := r.FormValue("trip-type")
		message := r.FormValue("message")

		to := "tess@twflyfishing.com"
		subject := fmt.Sprintf("New Website Inquiry from %s: %s", name, tripType)
		body := fmt.Sprintf("Name: %s\nEmail: %s\nTrip Interest: %s\n\nMessage:\n%s\n", name, email, tripType, message)

		msg := []byte(fmt.Sprintf("To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"%s\r\n", to, subject, body))

		// Try to send if smtp env vars are provided
		smtpHost := os.Getenv("SMTP_HOST")
		smtpPort := os.Getenv("SMTP_PORT")
		smtpUser := os.Getenv("SMTP_USER")
		smtpPass := os.Getenv("SMTP_PASS")

		if smtpHost != "" && smtpPort != "" {
			auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

			err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, []string{to}, msg)
			if err != nil {
				log.Printf("Failed to send email: %v", err)
				web.ContactError().Render(r.Context(), w)
				return
			}

			web.ContactSuccess().Render(r.Context(), w)
		}
	}
}
