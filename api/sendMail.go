package api

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"time"
)

func SendMail(body bytes.Buffer) {
	id := os.Getenv("GMAIL_ID")
	pass := os.Getenv("GMAIL_PASS")
	target := os.Getenv("TARGET_ID")

	auth := smtp.PlainAuth("", id, pass, "smtp.gmail.com")

	from := "dami@love.com"
	to := []string{target}

	aSubject := []string{"상 먹자", "부자 되거라", "오늘의 투자종목은?!!", "믿으셔야합니다.", "차 바꾸실 때 되셨죠?", "정석이 테슬라 개꿀맛"}
	timeSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(timeSource)
	nRand := random.Intn(100)
	fmt.Println(nRand)

	// 메시지 작성
	var szSubject bytes.Buffer
	szSubject.WriteString("Subject: ")
	szSubject.WriteString(aSubject[nRand%6])
	szSubject.WriteString("\n")
	headerSubject := szSubject.String()
	headerBlank := "\r\n"
	msg := []byte(headerSubject + headerBlank + body.String())

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	check(err)
}
