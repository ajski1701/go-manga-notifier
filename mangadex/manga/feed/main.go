package feed

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ajski1701/go-manga-notifier/mangadex/manga/title"
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

	fmt.Println(time.Now().Format(time.RFC3339), "Manga Feed Response status:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string([]byte(body)))
	var result FeedOutput
	json.Unmarshal([]byte(body), &result)
	//fmt.Println(string(body))

	for i := range result.Data {
		currentIndex := len(result.Data) - 1 - i
		currentItem := result.Data[currentIndex]
		language := currentItem.Attributes.TranslatedLanguage
		if language != "en" {
			continue
		}
		chapterNum := currentItem.Attributes.Chapter
		mangaTitle := ""
		chapterId := currentItem.Id
		chapterUrl := "https://mangadex.org/chapter/" + chapterId
		createdDate := currentItem.Attributes.CreatedAt

		parsedPublishDate, err := time.Parse(time.RFC3339, createdDate)

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
		localOutputMap["createdDate"] = parsedPublishDate.Format(time.RFC3339)
		outputArray = append(outputArray, localOutputMap)
	}
	return outputArray
}
