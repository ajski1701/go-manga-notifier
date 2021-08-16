package gomail

import (
	"strings"

	"gopkg.in/ini.v1"
)

func PrepMessageBody(manga map[string]string) string {
	//mangaTitle + " Chapter " + mangaChapter + '\n' + "" + chapterLink
	//title := manga["title"]
	chapter := manga["chapter"]
	url := manga["url"]
	//body := title + " Chapter " + chapter + "\n" + "" + url
	body := "Chapter " + chapter + "\n" + "" + url
	return body
}

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
