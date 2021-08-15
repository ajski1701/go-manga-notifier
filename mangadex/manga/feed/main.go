package feed

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mangadex-notifier/mangadex/manga/title"
	"net/http"
	"time"
)

func GetFollowedMangaFeedList(token string) []map[string]string {
	var outputArray []map[string]string
	var bearer = "Bearer " + token

	req, err := http.NewRequest("GET", "https://api.mangadex.org/user/follows/manga/feed?order[createdAt]=desc", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Manga Feed Response status:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string([]byte(body)))
	var result FeedOutput
	json.Unmarshal([]byte(body), &result)
	//fmt.Println(string(body))

	for i := range result.Results {
		currentIndex := len(result.Results) - 1 - i
		currentItem := result.Results[currentIndex]
		language := currentItem.Data.Attributes.TranslatedLanguage
		if language != "en" {
			continue
		}
		chapterNum := currentItem.Data.Attributes.Chapter
		mangaTitle := ""
		chapterId := currentItem.Data.Id
		chapterUrl := "https://mangadex.org/chapter/" + chapterId
		publishedDate := currentItem.Data.Attributes.PublishAt

		parsedPublishDate, err := time.Parse(time.RFC3339, publishedDate)

		if err != nil {
			panic(err)
		}

		for relationshipItem := range currentItem.Relationships {
			if currentItem.Relationships[relationshipItem].Type == "manga" {
				mangaTitle = title.GetTitle(currentItem.Relationships[relationshipItem].Id)
				break
			}
		}
		localOutputMap := map[string]string{}
		localOutputMap["title"] = mangaTitle
		localOutputMap["chapter"] = chapterNum
		localOutputMap["url"] = chapterUrl
		localOutputMap["publishedDate"] = parsedPublishDate.Format(time.RFC3339)
		//printOutput := "Title:" + mangaTitle + ": " + "\t\tChapter " + chapterNum + ". URL: " + chapterUrl
		outputArray = append(outputArray, localOutputMap)
		//fmt.Println(printOutput)
	}
	return outputArray
}
