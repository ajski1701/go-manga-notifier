package main

import (
	"fmt"
	"mangadex-notifier/config"
	"mangadex-notifier/gomail"
	"mangadex-notifier/mangadex/authentication"
	"mangadex-notifier/mangadex/manga/feed"
	"mangadex-notifier/mangadex/manga/title"
	"time"
)

func main() {
	user_cfg := config.LoadUserIni()
	app_cfg := config.LoadAppIni()
	to_email := gomail.ParseToEmail(user_cfg)
	sessionToken := authentication.GetAuth(user_cfg)
	manga := feed.GetFollowedMangaFeedList(sessionToken)
	lastRunTime := config.ParseLastRunTime(app_cfg)
	newLastRunTime := time.Now()

	for _, element := range manga {
		chapterCreationDate := title.ParseCreationDate(element["createdDate"])
		newLastRunTime = chapterCreationDate
		logTime := time.Now().Format(time.RFC3339)

		if lastRunTime.After(chapterCreationDate) || lastRunTime.Equal(chapterCreationDate) {
			fmt.Println(logTime, "Skipping alert for", element["title"], "Chapter", element["chapter"]+".")
			continue
		}

		emailBody := gomail.PrepMessageBody(element)
		alert, err := gomail.SendEmailSMTP(to_email, emailBody, element["title"], user_cfg)

		if err == nil && alert {
			fmt.Println(logTime, "Alert sent for", element["title"], "Chapter", element["chapter"]+".")
		} else {
			fmt.Println(logTime, "Failed to send alert for", element["title"], "Chapter", element["chapter"]+".")
		}
	}
	//Update the last run time ini
	config.UpdateAppIni(newLastRunTime)
}
