package main

import (
	"fmt"
	"net/smtp"
	"net/http"
)

func sendMail(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	subject := r.FormValue("subject")
	fromEmail := r.FormValue("from")
	text := r.FormValue("text")

	fmt.Println(name)

	smtpServer := "158.217.174.41"
	smtpPort := 25
	toEmail := "k646592@kansai-u.ac.jp"

	body := "From: " + name + "<" + fromEmail + ">\r\n" + "To: " + toEmail + "\r\n" + "Subject: " + subject + "\r\n\r\n" + text

	auth := smtp.PlainAuth("", "", "", smtpServer)

	tlsConfig := &tls.Config(
		InsecureSkipVerify: true,
		ServerName: smtpServer,
	)

	conn, err := tls.Dial("top", fmt.Sprintf("%s:%d", smtpServer, smtpPort), tlsConfig)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, fromEmail, []string{toEmail}, []byte(body))
	if err != nil {
		fmt.Println("Error sending mail:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, smtpServer)
	if err != nil {
		fmt.Println("Error creating SMTP client:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer client.Close()

	

	fmt.Println("Mail sent Successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Data received successfully"}`))
}

func main() {
	http.HandleFunc("/gomail", sendMail)

	fmt.Println("Server is running on :5000...")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}