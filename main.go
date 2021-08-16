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
	lastIndex := len(manga) - 1

	for i, element := range manga {
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

		//If we're in the last item in the array update the app ini to include the last publish date of the manga for next run
		if i == lastIndex {
			config.UpdateAppIni(chapterPublishDate)
		}
	}
}
