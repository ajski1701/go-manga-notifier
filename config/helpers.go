package config

import (
	"time"

	"gopkg.in/ini.v1"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseLastRunTime(cfg *ini.File) time.Time {
	lastRunTimeString := cfg.Section("").Key("last_chapter_publish_time").String()
	timeParsed, err := time.Parse(time.RFC3339, lastRunTimeString)
	if err != nil {
		panic(err)
	}
	return timeParsed
}
