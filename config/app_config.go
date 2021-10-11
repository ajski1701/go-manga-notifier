package config

import (
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/ini.v1"
)

const appIniTemplate = `last_app_run_time = `

const appIniName = "app_config.ini"

func LoadAppIni() *ini.File {
	cfg, err := ini.Load(appIniName)
	if err != nil {
		fmt.Printf("Fail to read file: %v\n", err)
		//Creates a ini file
		createAppIni()
		//Recursion on the fuction
		cfg = LoadAppIni()
	}
	return cfg
}

func createAppIni() {
	//Create ini
	fmt.Println("Creating template ini file.")
	err := ioutil.WriteFile(appIniName, []byte(appIniTemplate), 0644)
	check(err)
	cfg, err := ini.Load(appIniName)
	check(err)
	currentTime := time.Now()
	cfg.Section("").Key("last_app_run_time").SetValue(currentTime.Format(time.RFC3339))
	cfg.SaveTo(appIniName)
}

func UpdateAppIni() {
	cfg, err := ini.Load(appIniName)
	check(err)
	currentTime := time.Now()
	cfg.Section("").Key("last_app_run_time").SetValue(currentTime.Format(time.RFC3339))
	cfg.SaveTo(appIniName)
}
