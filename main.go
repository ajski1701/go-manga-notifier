package main

import (
	"fmt"
	"mangadex-notifier/config"
	"mangadex-notifier/gomail"
	"mangadex-notifier/mangadex/authentication"
	"mangadex-notifier/mangadex/manga/feed"
	"mangadex-notifier/mangadex/manga/title"
)

func main() {
	user_cfg := config.LoadUserIni()
	app_cfg := config.LoadAppIni()
	to_email := gomail.ParseToEmail(user_cfg)
	sessionToken := authentication.GetAuth(user_cfg)
	manga := feed.GetFollowedMangaFeedList(sessionToken)
	lastRunTime := config.ParseLastRunTime(app_cfg)

	for _, element := range manga {
		chapterPublishDate := title.ParsePublishDate(element["publishedDate"])
		if lastRunTime.After(chapterPublishDate) {
			fmt.Println("Skipping alert for " + element["title"] + " Chapter " + element["chapter"] + ".")
			continue
		}

		emailBody := gomail.PrepMessageBody(element)
		alert, err := gomail.SendEmailSMTP(to_email, emailBody, element["title"], user_cfg)

		if err == nil && alert {
			fmt.Println("Alert sent for " + element["title"] + " Chapter " + element["chapter"] + ".")
		} else {
			fmt.Println("Failed to send alert for " + element["title"] + " Chapter " + element["chapter"] + ".")
		}
	}
	//Update the last run time ini
	config.UpdateAppIni()
}
