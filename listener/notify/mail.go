package notify

import (
	"errors"
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/simon-liu/supervisor-event-listener/event"
	"strings"
)

type Mail struct{}

func (mail *Mail) Send(message event.Message) error {
	body := message.String()
	body = strings.Replace(body, "\n", "<br>", -1)
	gomailMessage := gomail.NewMessage()
	gomailMessage.SetHeader("From", Conf.MailServer.From)
	gomailMessage.SetHeader("To", Conf.MailUser.Email...)
	gomailMessage.SetHeader("Subject", "Supervisor事件通知")
	gomailMessage.SetBody("text/html", body)
	mailer := gomail.NewDialer(
		Conf.MailServer.Host,
		Conf.MailServer.Port,
		Conf.MailServer.Login,
		Conf.MailServer.Password,
	)
	err := mailer.DialAndSend(gomailMessage)
	if err == nil {
		return nil
	}
	errorMessage := fmt.Sprintf("邮件发送失败#%s", err.Error())

	return errors.New(errorMessage)
}
