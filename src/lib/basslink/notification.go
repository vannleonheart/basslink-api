package basslink

import (
	"bytes"
	"html/template"
)

type EmailNotificationMesage struct {
	To       string
	Subject  string
	Template string
	Data     map[string]interface{}
}

func (app *App) HandleEmailNotification(message *EmailNotificationMesage) {
	var tpl Template

	if err := app.DB.Connection.Where("id = ?", message.Template).First(&tpl).Error; err != nil {
		return
	}

	html := ""

	if templateHtml, parseError := template.New("email").Parse(tpl.Data); parseError == nil {
		var buf bytes.Buffer

		if err := templateHtml.Execute(&buf, message.Data); err == nil {
			html = buf.String()
		}
	}

	if _, err := app.Mailgun.SendEmail(message.To, message.Subject, &html); err != nil {
		return
	}

	return
}
