package handler

import (
	"fmt"
	"net/smtp"
	"time"
)

func email_send(email string, name string, dateTime time.Time, category string) {

	smtpHost := "smtp.yandex.ru"
	smtpPort := "587"

	from := "**********"
	password := "**********"

	to := []string{email}

	formattedDate := dateTime.Format("02.01.2006 15:04")

	subject := "Subject: Уведомление о мероприятии\r\n"
	toHeader := "To: " + email + "\r\n"
	fromHeader := "From: Event System <" + from + ">\r\n"

	mime := "MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n"

	body := fmt.Sprintf(`
<html>
<body>
	<h2>Привет!</h2>
	<p>Назначено новое мероприятие:</p>

	<ul>
		<li><b>Название:</b> %s</li>
		<li><b>Дата:</b> %s</li>
		<li><b>Категория:</b> %s</li>
	</ul>
</body>
</html>
`, name, formattedDate, category)

	msg := subject + fromHeader + toHeader + mime + body

	auth := smtp.PlainAuth("", from, password, smtpHost)

	fmt.Println("Connecting to SMTP...")

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println("❌ SMTP ERROR:", err)
		return
	}

}
