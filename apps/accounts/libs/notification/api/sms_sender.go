package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"backend/libs/configuration"
	log "backend/libs/logger"
)

type SmsSender struct {
	User     string
	Password string
	From     string
}

func NewSmsSender(user, passwd string) SmsSender {
	return SmsSender{
		User:     user,
		Password: passwd,
	}
}

func (sender SmsSender) SendMessage(phoneNumber string, code string) (map[string]interface{}, error) {
	log.Info("Solicitando envio de sms para o número:", phoneNumber)

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + sender.User + "/Messages.json"

	msgData := url.Values{}
	msgData.Set("To", phoneNumber)
	msgData.Set("From", configuration.Get().GetEnvConfString("sms.from"))

	body := fmt.Sprintf("Use %s como código de acesso para Estratégia em Áudio", code)
	msgData.Set("Body", body)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(sender.User, sender.Password)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	var data map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Error(err)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		infoMessage := fmt.Sprintf("SMS enviado para o numero %s, SID: %s", phoneNumber, data["sid"])
		log.Info(infoMessage)

		return data, nil
	} else {
		errorMessage := fmt.Sprintf("%.0f - %s", data["code"], data["message"])

		log.Error("Error ao solicitar envio de sms:", errorMessage)
		return data, errors.New("Erro ao solicitar SMS")
	}
}
