package title

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ajski1701/go-manga-notifier/mangadex/models"
)

func ParseCreationDate(dateStr string) time.Time {
	timeParsed, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		panic(err)
	}
	return timeParsed
}

func GetTitle(mangaId string) string {
	url := "https://api.mangadex.org/manga/" + mangaId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result models.MangaOutput
	json.Unmarshal([]byte(body), &result)
	titleLanguage := ""
	for element := range result.Data.Attributes.Title {
		titleLanguage = element
	}
	return result.Data.Attributes.Title[titleLanguage]
}
