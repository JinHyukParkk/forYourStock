package api

import (
	"bytes"
	"net/smtp"
	"os"
)

func SendMail(body bytes.Buffer) {
	id := os.Getenv("GMAIL_ID")
	pass := os.Getenv("GMAIL_PASS")
	target := os.Getenv("TARGET_ID")

	auth := smtp.PlainAuth("", id, pass, "smtp.gmail.com")

	from := "dami@love.com"
	to := []string{target}

	// 메시지 작성
	headerSubject := "Subject: 부자 되거라\r\n"
	headerBlank := "\r\n"
	msg := []byte(headerSubject + headerBlank + body.String())

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	check(err)
}
