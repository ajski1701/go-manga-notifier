package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/ini.v1"
)

const userIniTemplate = `[email]
from = <from_gmail_email_address>
password = <gmail_password>
to = <comma_delimited_recepient_emails>

[mangadex]
username = <mangadex_username>
password = <mangadex_password>`

const userIniName = "user_config.ini"

func LoadUserIni() *ini.File {
	cfg, err := ini.Load(userIniName)
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		createUserIni()
		os.Exit(1)
	}
	//fmt.Println("email user:", cfg.Section("email").Key("from").String())
	//fmt.Println("mangadex user:", cfg.Section("mangadex").Key("username").String())
	return cfg
}

func createUserIni() {
	fmt.Println("Creating template ini file.")
	err := ioutil.WriteFile(userIniName, []byte(userIniTemplate), 0644)
	check(err)
}
