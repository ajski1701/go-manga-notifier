package parse

import (
	"strings"
	"time"

	"gopkg.in/ini.v1"
)

func ParseToEmail(cfg *ini.File) []string {
	var output []string
	string := cfg.Section("email").Key("to").String()
	string = strings.ReplaceAll(string, " ", "")
	output = strings.Split(string, ",")
	for i, element := range output {
		element = strings.TrimSpace(element)
		output[i] = element
	}
	return output
}

func ParseLastRunTime(cfg *ini.File) time.Time {
	lastRunTimeString := cfg.Section("").Key("last_chapter_publish_time").String()
	timeParsed, err := time.Parse(time.RFC3339, lastRunTimeString)
	if err != nil {
		panic(err)
	}
	return timeParsed
}

func ParsePublishDate(dateStr string) time.Time {
	timeParsed, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		panic(err)
	}
	return timeParsed
}
