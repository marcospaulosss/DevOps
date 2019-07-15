package notification

import (
	"backend/apps/accounts/libs/notification/api"
	"backend/libs/configuration"
	log "backend/libs/logger"
)

type NotificationInterface interface {
	SendEmailNotification(email string, code string)
	SendSmsNotification(phone string, code string) (string, error)
}

type Notification struct {
	emailSender api.EmailSender
	smsSender   api.SmsSender
}

func NewNotification() NotificationInterface {
	config := configuration.Get()
	smtpUser := config.GetEnvConfString("smtp.user")
	smtpPasswd := config.GetEnvConfString("smtp.secret")

	smsID := config.GetEnvConfString("sms.id")
	smsToken := config.GetEnvConfString("sms.secret")
	return &Notification{
		emailSender: api.NewEmailSender(smtpUser, smtpPasswd),
		smsSender:   api.NewSmsSender(smsID, smsToken),
	}
}

func (this *Notification) SendEmailNotification(emailTo string, code string) {
	subject := "Chegou seu Código de verificação do Email"
	message := "Código de validação do email: " + code

	log.Info("Preparando o corpo do email")
	bodyMessage := this.emailSender.WritePlainEmail(emailTo, subject, message)

	log.Info("Enviando email para", emailTo)
	this.emailSender.SendMail(emailTo, subject, bodyMessage)
}

func (this *Notification) SendSmsNotification(phone string, code string) (string, error) {
	phone = "+" + phone
	resp, err := this.smsSender.SendMessage(phone, code)
	if err != nil {
		return "", err
	}

	return resp["sid"].(string), nil
}
