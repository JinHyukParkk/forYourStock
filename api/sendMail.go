package api

import (
	"net/smtp"
	"os"
)

func SendMail() {
	id := os.Getenv("GMAIL_ID")
	pass := os.Getenv("GMAIL_PASS")
	target := os.Getenv("TARGET_ID")

	auth := smtp.PlainAuth("", id, pass, "smtp.gmail.com")

	from := "dami@love.com"
	to := []string{target}

	// 메시지 작성
	headerSubject := "Subject: 다미야 사랑해\r\n"
	headerBlank := "\r\n"
	body := "테스트다"
	msg := []byte(headerSubject + headerBlank + body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	check(err)
}
