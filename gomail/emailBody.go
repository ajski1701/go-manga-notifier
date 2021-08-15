package gomail

func PrepMessageBody(manga map[string]string) string {
	//mangaTitle + " Chapter " + mangaChapter + '\n' + "" + chapterLink
	//title := manga["title"]
	chapter := manga["chapter"]
	url := manga["url"]
	//body := title + " Chapter " + chapter + "\n" + "" + url
	body := "Chapter " + chapter + "\n" + "" + url
	return body
}
