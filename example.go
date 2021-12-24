package main

import (
	"fmt"
	"github.com/huzhengyao0/send-email-ssl/ssl"
)

func main() {
	mailUtils := ssl.EMailUtils{
		Host:       "example.com",
		ServerAddr: "example.com:port",
		Username:   "example",
		Password:   "example",
	}
	// send mail
	err := mailUtils.SendMail("example@example.com", "aa@example.com, bb@example.com", "TEST MAIL", "TEST MAIL HTML", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}