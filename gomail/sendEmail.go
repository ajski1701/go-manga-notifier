//https://medium.com/wesionary-team/sending-emails-with-go-golang-using-smtp-gmail-and-oauth2-185ee12ab306

package gomail

import (
	"fmt"
	"net/smtp"
	"strings"

	"gopkg.in/ini.v1"
)

var emailAuth smtp.Auth

func SendEmailSMTP(to []string, emailBody string, subject string, cfg *ini.File) (bool, error) {
	emailHost := "smtp.gmail.com"
	emailFrom := cfg.Section("email").Key("from").String()
	emailPassword := cfg.Section("email").Key("password").String()
	emailPort := 587

	emailAuth = smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	fullBody := []byte("From: " + emailFrom + "\nTo: " + (strings.Join(to[:], ", ")) + "\nSubject:" + subject + "\n" + emailBody + "\n")
	addr := fmt.Sprintf("%v:%v", emailHost, emailPort)

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, fullBody); err != nil {
		return false, err
	}
	return true, nil
}
