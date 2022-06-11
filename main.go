package main

import (
	"fmt"
	"time"

	"github.com/ajski1701/go-manga-notifier/config"
	"github.com/ajski1701/go-manga-notifier/gomail"
	"github.com/ajski1701/go-manga-notifier/mangadex/authentication"
	"github.com/ajski1701/go-manga-notifier/mangadex/manga/feed"
	"github.com/ajski1701/go-manga-notifier/mangadex/manga/title"
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
