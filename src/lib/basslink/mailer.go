package basslink

import (
	"fmt"
	mj "github.com/mailjet/mailjet-apiv3-go"
	"strconv"
)

const (
	MailerServiceMailjet = "mailjet"

	MailjetPostRegistrationTemplateId   = "6591635"
	MailjetRecoverAccountTemplateId     = "6591645"
	MailjetSubagentInvitationTemplateId = "6606077"
)

type Mailer struct {
	Config   MailerConfig
	services map[string]interface{}
}

type MailerConfig struct {
	Sender struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"sender"`
	Endpoints map[string]string                 `json:"endpoints"`
	Services  map[string]map[string]interface{} `json:"services"`
}

func NewMailer(config *MailerConfig) *Mailer {
	ml := Mailer{
		Config:   *config,
		services: make(map[string]interface{}),
	}

	if config.Services != nil {
		for service, params := range config.Services {
			ml.services[service] = createMailService(service, params)
		}
	}

	return &ml
}

func (m *Mailer) SendMailTemplate(service, to, subject, templateId string, templateData map[string]interface{}) (interface{}, error) {
	serviceI, ok := m.services[service]
	if !ok || serviceI == nil {
		return nil, fmt.Errorf("mailer service %s not found", service)
	}

	switch service {
	case MailerServiceMailjet:
		return m.sendMailTemplateMailjet(serviceI.(*mj.Client), to, subject, templateId, templateData)
	}

	return nil, nil
}

func (m *Mailer) sendMailTemplateMailjet(client *mj.Client, to, subject, templateId string, templateData map[string]interface{}) (interface{}, error) {
	intTemplateid, _ := strconv.ParseInt(templateId, 10, 64)
	messagesInfo := []mj.InfoMessagesV31{
		{
			From: &mj.RecipientV31{
				Email: m.Config.Sender.Email,
				Name:  m.Config.Sender.Name,
			},
			To: &mj.RecipientsV31{
				mj.RecipientV31{
					Email: to,
					Name:  "Test",
				},
			},
			Subject:          subject,
			TemplateID:       intTemplateid,
			TemplateLanguage: true,
			Variables:        templateData,
		},
	}
	messages := mj.MessagesV31{Info: messagesInfo}
	res, err := client.SendMailV31(&messages)
	if err != nil {
		return nil, err
	}
	return res, err
}

func createMailService(service string, params map[string]interface{}) interface{} {
	switch service {
	case MailerServiceMailjet:
		apiKey, secretKey := "", ""
		if apiKeyI, ok := params["api_key"]; ok {
			apiKey = apiKeyI.(string)
		}
		if secretKeyI, ok := params["secret_key"]; ok {
			secretKey = secretKeyI.(string)
		}
		if len(apiKey) > 0 && len(secretKey) > 0 {
			return mj.NewMailjetClient(apiKey, secretKey)
		}
	}

	return nil
}
