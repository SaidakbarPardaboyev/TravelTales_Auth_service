// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"html/template"
// 	"net/smtp"
// )

// func main() {

// 	// client := redis.NewClient(&redis.Options{
// 	// 	Addr:     "localhost:6379",
// 	// 	Password: "", // no password set
// 	// 	DB:       0,  // use default DB
// 	// })

// 	code := "123456"

// 	// err := client.Set(context.Background(), "Key-test", code, time.Minute*5).Err()
// 	// if err != nil {
// 	// 	return
// 	// }

// 	// result, err := client.Get(context.Background(), "Key-test").Result()
// 	// if err != nil {
// 	// 	return
// 	// }

// 	// println(result)

// 	SendCode("pardaboyevsaidakbar103@gmail.com", code)
// }

// func SendCode(email string, code string) {
// 	// sender data
// 	from := "kupalovv.muhammadjon@gmail.com"
// 	password := "vump lxbf awbv slck"

// 	// Receiver email address
// 	to := []string{
// 		email,
// 	}

// 	// smtp server configuration.
// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"

// 	// Authentication.
// 	auth := smtp.PlainAuth("", from, password, smtpHost)

// 	t, _ := template.ParseFiles("template.html")

// 	var body bytes.Buffer

// 	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// 	body.Write([]byte(fmt.Sprintf("Subject: Your verification code \n%s\n\n", mimeHeaders)))

// 	t.Execute(&body, struct {
// 		Passwd string
// 	}{

// 		Passwd: code,
// 	})

// 	// Sending email.
// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Email sended to:", email)

// }

package main

import (
	"fmt"
	"net/smtp"
)

func sendEmail(to, subject, body string) error {
	// Set up authentication information.
	from := "kupalovv.muhammadjon@gmail.com"
	password := "vump lxbf awbv slck"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Set up email content.
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	// Send the email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func main() {
	to := "pardaboyevsaidakbar103@gmail.com"
	subject := "Test Email"
	body := "Hello, this is a test email from Go!"

	err := sendEmail(to, subject, body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
