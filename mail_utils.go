package mail

import (
	"fmt"
	"net/smtp"
	"strings"
)

type EMailUtils struct {
	Host       string `json:"host"`
	ServerAddr string `json:"serverAddr"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func NewMailUtils(host, username, password, serverAddr string) *EMailUtils {
	return &EMailUtils{
		Host:       host,
		ServerAddr: serverAddr,
		Username:   username,
		Password:   password,
	}
}

/**
sender
receivers  eg aaa@example.aa,bb@example.bb
subject
content
attachments
*/
func (mailUtils EMailUtils) SendMail(sender, receivers, subject, content string, attachments []string) error {
	email := NewEmail()
	email.From = fmt.Sprintf("%s<%s>", strings.Split(sender, "@")[0], sender)
	email.To = strings.Split(receivers, ",")
	email.Subject = subject
	email.HTML = []byte(content)
	attas := make([]*Attachment, len(attachments))
	for _, attachment := range attachments {
		atta, err := email.AttachFile(attachment)
		if err != nil {
			return err
		}
		attas = append(attas, atta)
	}
	return email.Send(mailUtils.ServerAddr, smtp.PlainAuth("", mailUtils.Username, mailUtils.Password, mailUtils.Host))
}
