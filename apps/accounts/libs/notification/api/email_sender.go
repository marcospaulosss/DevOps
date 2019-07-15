package api

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"

	"backend/libs/configuration"
	log "backend/libs/logger"
)

type EmailSender struct {
	User     string
	Password string
	From     string
}

func NewEmailSender(User, Password string) EmailSender {
	sender := EmailSender{}
	sender.User = User
	sender.Password = Password
	sender.From = configuration.Get().GetEnvConfString("smtp.sender")
	return sender
}

func (sender EmailSender) SendMail(dest string, subject, bodyMessage string) {
	config := configuration.Get()
	var host = config.GetEnvConfString("smtp.server")
	var port = config.GetEnvConfString("smtp.port")
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	servername := fmt.Sprintf("%s:%s", host, port)

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Error("Error connecting to mail server.", err)
		return
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Error("Error creating mail client.", err)
		return
	}

	auth := smtp.PlainAuth("", sender.User, sender.Password, host)
	if err = c.Auth(auth); err != nil {
		log.Error("Error on mail authentication.", err)
		return
	}

	if err = c.Mail(sender.From); err != nil {
		log.Error("Error in From.", err)
		return
	}

	if err = c.Rcpt(dest); err != nil {
		log.Error("Error in To", dest, err)
		return
	}

	w, err := c.Data()
	if err != nil {
		log.Error("Error in Data.", err)
		return
	}

	_, err = w.Write([]byte(bodyMessage))
	if err != nil {
		log.Error("Error writing data.", err)
		return
	}

	err = w.Close()
	if err != nil {
		log.Error("Error closing.", err)
		return
	}

	c.Quit()

	log.Info("Email enviado com sucesso!", subject)
}

func (sender EmailSender) WriteEmail(to string, contentType, subject, bodyMessage string) string {
	header := make(map[string]string)
	header["From"] = sender.From

	header["To"] = to
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(bodyMessage))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}

func (sender *EmailSender) WriteHTMLEmail(to string, subject, bodyMessage string) string {
	return sender.WriteEmail(to, "text/html", subject, bodyMessage)
}

func (sender *EmailSender) WritePlainEmail(to string, subject, bodyMessage string) string {
	return sender.WriteEmail(to, "text/plain", subject, bodyMessage)
}
