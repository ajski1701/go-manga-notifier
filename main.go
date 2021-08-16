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

	for i, element := range manga {
		chapterPublishDate := title.ParsePublishDate(element["publishedDate"])
		if lastRunTime.After(chapterPublishDate) {
			fmt.Println(element["title"] + " Chapter " + element["chapter"] + " is not new. Skipping.")
			continue
		}

		emailBody := gomail.PrepMessageBody(element)
		fmt.Println(gomail.SendEmailSMTP(to_email, emailBody, element["title"], user_cfg))

		//If we're in the last item in the array update the app ini to include the last publish date of the manga
		if i == len(manga)-1 {
			config.UpdateAppIni(chapterPublishDate)
		}
	}
}
